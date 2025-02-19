// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package emr

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/emr"
	"github.com/aws/aws-sdk-go/service/emr/emriface"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// []*SERVICE.Tag handling

// Tags returns emr service tags.
func Tags(tags tftags.KeyValueTags) []*emr.Tag {
	result := make([]*emr.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &emr.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from emr service tags.
func KeyValueTags(ctx context.Context, tags []*emr.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// GetTagsIn returns emr service tags from Context.
// nil is returned if there are no input tags.
func GetTagsIn(ctx context.Context) []*emr.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// SetTagsOut sets emr service tags in Context.
func SetTagsOut(ctx context.Context, tags []*emr.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}

// UpdateTags updates emr service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.

func UpdateTags(ctx context.Context, conn emriface.EMRAPI, identifier string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	if removedTags := oldTags.Removed(newTags); len(removedTags) > 0 {
		input := &emr.RemoveTagsInput{
			ResourceId: aws.String(identifier),
			TagKeys:    aws.StringSlice(removedTags.IgnoreSystem(names.EMR).Keys()),
		}

		_, err := conn.RemoveTagsWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	if updatedTags := oldTags.Updated(newTags); len(updatedTags) > 0 {
		input := &emr.AddTagsInput{
			ResourceId: aws.String(identifier),
			Tags:       Tags(updatedTags.IgnoreSystem(names.EMR)),
		}

		_, err := conn.AddTagsWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// UpdateTags updates emr service tags.
// It is called from outside this package.
func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return UpdateTags(ctx, meta.(*conns.AWSClient).EMRConn(), identifier, oldTags, newTags)
}
