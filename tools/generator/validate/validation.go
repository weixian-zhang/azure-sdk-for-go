package validate

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Azure/azure-sdk-for-go/tools/generator/autorest/model"
	"github.com/Azure/azure-sdk-for-go/tools/generator/utils"
)

// MetadataValidateFunc is a function that validates a metadata is legal or not
type MetadataValidateFunc func(ctx *MetadataValidateContext, tag string, metadata model.Metadata) error

// MetadataValidateContext describes the context needed in validation of the metadata
type MetadataValidateContext struct {
	Readme     string
	SDKRoot    string
	Validators []MetadataValidateFunc
}

// Validate validates the metadata
func (ctx *MetadataValidateContext) Validate(tag string, metadata model.Metadata) []error {
	var errors []error
	for _, validator := range ctx.Validators {
		if err := validator(ctx, tag, metadata); err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

func (ctx *MetadataValidateContext) getRelPackagePath(pkgPath string) (string, error) {
	rel, err := filepath.Rel(ctx.SDKRoot, pkgPath)
	if err != nil {
		return "", fmt.Errorf("cannot get relative path from output-folder '%s' to the root directory '%s': %+v", pkgPath, ctx.SDKRoot, err)
	}
	return utils.NormalizePath(rel), nil
}

func rootCheck(ctx *MetadataValidateContext, metadata model.Metadata) error {
	r := filepath.Clean(ctx.SDKRoot)
	o := filepath.Clean(metadata.PackagePath())
	if !strings.HasPrefix(o, r) {
		return fmt.Errorf("the output-folder '%s' is not under root directory '%s', the output-folder is either not configured or not correctly configured", metadata.PackagePath(), ctx.SDKRoot)
	}
	return nil
}

// PreviewCheck ensures the output-folder of a preview package is under the preview sub-directory
func PreviewCheck(ctx *MetadataValidateContext, tag string, metadata model.Metadata) error {
	if err := rootCheck(ctx, metadata); err != nil {
		return err
	}
	rel, err := ctx.getRelPackagePath(metadata.PackagePath())
	if err != nil {
		return err
	}
	if isPreviewPackage(metadata.SwaggerFiles()) {
		if !previewOutputRegex.MatchString(rel) {
			return fmt.Errorf("the output-folder of a preview package '%s' must be under the `preview` subdirectory", tag)
		}
	} else {
		if previewOutputRegex.MatchString(rel) {
			return fmt.Errorf("the output-folder of a stable package '%s' must NOT be under the `preview` subdirectory", tag)
		}
	}
	return nil
}

// MgmtCheck ensures that the management plane package has the correct output-folder
func MgmtCheck(ctx *MetadataValidateContext, tag string, metadata model.Metadata) error {
	if isMgmtPackage(ctx.Readme) {
		rel, err := ctx.getRelPackagePath(metadata.PackagePath())
		if err != nil {
			return err
		}
		if !mgmtOutputRegex.MatchString(rel) {
			return fmt.Errorf("the output-folder of a management plane package '%s' must be in this pattern: '^services(/preview)?/[^/]+/mgmt/[^/]+/[^/]+$'", tag)
		}
	}
	return nil
}

func isPreviewPackage(inputFiles []string) bool {
	for _, inputFile := range inputFiles {
		if isPreviewSwagger(inputFile) {
			return true
		}
	}
	return false
}

func isPreviewSwagger(inputFile string) bool {
	return previewSwaggerRegex.MatchString(inputFile)
}

func isMgmtPackage(readme string) bool {
	return mgmtReadmeRegex.MatchString(readme)
}

var (
	previewSwaggerRegex = regexp.MustCompile(`^preview|.+[/\\]preview[/\\]`)
	previewOutputRegex  = regexp.MustCompile(`^services/preview/`)
	mgmtReadmeRegex     = regexp.MustCompile(`[/\\]resource-manager[/\\]`)
	mgmtOutputRegex     = regexp.MustCompile(`^services(/preview)?/[^/]+/mgmt/[^/]+/[^/]+$`)
)