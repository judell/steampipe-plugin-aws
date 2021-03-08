package aws

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
	pb "github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsKinesisStream(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_kinesis_stream",
		Description: "AWS Kinesis Stream",
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("stream_name"),
			ShouldIgnoreError: isNotFoundError([]string{"ResourceNotFoundException", "InvalidParameter"}),
			Hydrate:           describeStream,
		},
		List: &plugin.ListConfig{
			Hydrate: listStreams,
		},
		GetMatrixItem: BuildRegionList,
		Columns: awsRegionalColumns([]*plugin.Column{
			{
				Name:        "stream_name",
				Description: "The name of the stream being described.",
				Type:        pb.ColumnType_STRING,
				Transform:   transform.FromField("StreamDescription.StreamName"),
			},
			{
				Name:        "stream_arn",
				Description: "The Amazon Resource Name (ARN) for the stream being described.",
				Type:        pb.ColumnType_STRING,
				Hydrate:     describeStream,
				Transform:   transform.FromField("StreamDescription.StreamARN"),
			},
			{
				Name:        "stream_status",
				Description: "The current status of the stream being described.",
				Type:        pb.ColumnType_STRING,
				Hydrate:     describeStream,
				Transform:   transform.FromField("StreamDescription.StreamStatus"),
			},
			{
				Name:        "stream_creation_timestamp",
				Description: "The approximate time that the stream was created.",
				Type:        pb.ColumnType_TIMESTAMP,
				Hydrate:     describeStream,
				Transform:   transform.FromField("StreamDescription.StreamCreationTimestamp"),
			},
			{
				Name:        "encryption_type",
				Description: "The server-side encryption type used on the stream.",
				Type:        pb.ColumnType_STRING,
				Hydrate:     describeStream,
				Transform:   transform.FromField("StreamDescription.EncryptionType"),
			},
			{
				Name:        "key_id",
				Description: "The GUID for the customer-managed AWS KMS key to use for encryption.",
				Type:        pb.ColumnType_STRING,
				Hydrate:     describeStream,
				Transform:   transform.FromField("StreamDescription.KeyId"),
			},
			{
				Name:        "retention_period_hours",
				Description: "The current retention period, in hours.",
				Type:        pb.ColumnType_INT,
				Hydrate:     describeStream,
				Transform:   transform.FromField("StreamDescription.RetentionPeriodHours"),
			},
			{
				Name:        "consumer_count",
				Description: "The number of enhanced fan-out consumers registered with the stream.",
				Type:        pb.ColumnType_INT,
				Hydrate:     describeStreamSummary,
				Transform:   transform.FromField("StreamDescriptionSummary.ConsumerCount"),
			},
			{
				Name:        "open_shard_count",
				Description: "The number of open shards in the stream.",
				Type:        pb.ColumnType_INT,
				Hydrate:     describeStreamSummary,
				Transform:   transform.FromField("StreamDescriptionSummary.OpenShardCount"),
			},
			{
				Name:        "has_more_shards",
				Description: "If set to true, more shards in the stream are available to describe.",
				Type:        pb.ColumnType_BOOL,
				Hydrate:     describeStream,
				Transform:   transform.FromField("StreamDescription.HasMoreShards"),
			},
			{
				Name:        "shards",
				Description: "The shards that comprise the stream.",
				Type:        pb.ColumnType_JSON,
				Hydrate:     describeStream,
				Transform:   transform.FromField("StreamDescription.Shards"),
			},
			{
				Name:        "enhanced_monitoring",
				Description: "Represents the current enhanced monitoring settings of the stream.",
				Type:        pb.ColumnType_JSON,
				Hydrate:     describeStream,
				Transform:   transform.FromField("StreamDescription.EnhancedMonitoring"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with StreamName.",
				Type:        pb.ColumnType_JSON,
				Hydrate:     getAwsKinesisStreamTags,
				Transform:   transform.FromField("Tags"),
			},
			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        pb.ColumnType_STRING,
				Transform:   transform.FromField("StreamDescription.StreamName"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        pb.ColumnType_JSON,
				Hydrate:     getAwsKinesisStreamTags,
				Transform:   transform.FromField("Tags").Transform(kinesisTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        pb.ColumnType_JSON,
				Hydrate:     getAwsKinesisStreamAkas,
				Transform:   transform.FromValue(),
			},
		}),
	}
}

//// LIST FUNCTION

func listStreams(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO put me in helper function
	var region string
	matrixRegion := plugin.GetMatrixItem(ctx)[matrixKeyRegion]
	if matrixRegion != nil {
		region = matrixRegion.(string)
	}
	plugin.Logger(ctx).Trace("listStreams", "AWS_REGION", region)

	// Create session
	svc, err := KinesisService(ctx, d, region)
	if err != nil {
		return nil, err
	}

	// List call
	err = svc.ListStreamsPages(
		&kinesis.ListStreamsInput{},
		func(page *kinesis.ListStreamsOutput, isLast bool) bool {
			for _, streams := range page.StreamNames {
				d.StreamListItem(ctx, &kinesis.DescribeStreamOutput{
					StreamDescription: &kinesis.StreamDescription{
						StreamName: streams,
					},
				})
			}
			return *page.HasMoreStreams
		},
	)

	return nil, err
}

//// HYDRATE FUNCTIONS

func describeStream(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("describeStream")

	// TODO put me in helper function
	var region string
	matrixRegion := plugin.GetMatrixItem(ctx)[matrixKeyRegion]
	if matrixRegion != nil {
		region = matrixRegion.(string)
	}
	var streamName string
	if h.Item != nil {
		streamName = *h.Item.(*kinesis.DescribeStreamOutput).StreamDescription.StreamName
	} else {
		quals := d.KeyColumnQuals
		streamName = quals["stream_name"].GetStringValue()
	}

	// get service
	svc, err := KinesisService(ctx, d, region)
	if err != nil {
		return nil, err
	}

	// Build the params
	params := &kinesis.DescribeStreamInput{
		StreamName: aws.String(streamName),
	}

	// Get call
	data, err := svc.DescribeStream(params)
	if err != nil {
		logger.Debug("describeStream__", "ERROR", err)
		return nil, err
	}
	return data, nil
}

// API call for Stream Summary
func describeStreamSummary(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("describeStreamSummary")
	
	// TODO put me in helper function
	var region string
	matrixRegion := plugin.GetMatrixItem(ctx)[matrixKeyRegion]
	if matrixRegion != nil {
		region = matrixRegion.(string)
	}
	var streamName string
		streamName = *h.Item.(*kinesis.DescribeStreamOutput).StreamDescription.StreamName

	// get service
	svc, err := KinesisService(ctx, d, region)
	if err != nil {
		return nil, err
	}

	// Build the params
	params := &kinesis.DescribeStreamSummaryInput{
		StreamName: aws.String(streamName),
	}

	// Get call
	data, err := svc.DescribeStreamSummary(params)
	if err != nil {
		logger.Debug("describeStreamSummary__", "ERROR", err)
		return nil, err
	}
	return data, nil
}

// API call for fetching tag list
func getAwsKinesisStreamTags(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("getAwsKinesisStreamTags")

	// TODO put me in helper function
	var region string
	matrixRegion := plugin.GetMatrixItem(ctx)[matrixKeyRegion]
	if matrixRegion != nil {
		region = matrixRegion.(string)
	}

	streamName := *h.Item.(*kinesis.DescribeStreamOutput).StreamDescription.StreamName

	// Create Session
	svc, err := KinesisService(ctx, d, region)
	if err != nil {
		return nil, err
	}

	// Build the params
	params := &kinesis.ListTagsForStreamInput{
		StreamName: &streamName,
	}

	// Get call
	op, err := svc.ListTagsForStream(params)
	if err != nil {
		logger.Debug("getAwsKinesisStreamTags", "ERROR", err)
		return nil, err
	}

	return op, nil
}

func getAwsKinesisStreamAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getAwsKinesisStreamAkas")
	streamName := *h.Item.(*kinesis.DescribeStreamOutput).StreamDescription.StreamName
	c, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := c.(*awsCommonColumnData)
	aka := "arn:" + commonColumnData.Partition + ":kinesis:" + commonColumnData.Region + ":" + commonColumnData.AccountId + ":stream"

	if strings.HasPrefix(streamName, "/") {
		aka = aka + streamName
	} else {
		aka = aka + "/" + streamName
	}

	return []string{aka}, nil
}

//// TRANSFORM FUNCTIONS

func kinesisTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("ssmTagListToTurbotTags")
	tagList := d.Value.([]*kinesis.Tag)

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
