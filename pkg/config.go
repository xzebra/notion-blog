package notion_blog

type BlogConfig struct {
	DatabaseID string `mapstructure:"DatabaseID" usage:"ID of the Notion database of the blog."`

	ImagesLink    string `mapstructure:"ImagesLink" usage:"Directory in which the hugo content will be generated."`
	ImagesFolder  string `mapstructure:"ImagesFolder" usage:"Directory in which the static images will be stored. E.g.: ./web/static/images"`
	ContentFolder string `mapstructure:"ContentFolder" usage:"URL beggining to link the static images. E.g.: /images"`
	ArchetypeFile string `mapstructure:"ArchetypeFile" usage:"Route to the archetype file to generate the header."`

	// Optional:
	PropertyTitle       string `mapstructure:"property_title" usage:"Title property name in Notion."`
	PropertyDescription string `mapstructure:"PropertyDescription" usage:"Description property name in Notion."`
	PropertyTags        string `mapstructure:"PropertyTags" yaml:"property_tags" usage:"Tags multi-select porperty name in Notion."`
	PropertyCategories  string `mapstructure:"PropertyCategories" usage:"Categories multi-select porperty name in Notion."`

	FilterProp     string   `mapstructure:"FilterProp" usage:"Property of the filter to apply to a select value of the articles."`
	FilterValue    []string `mapstructure:"FilterValue" usage:"Value of the filter to apply to the Notion articles database."`
	PublishedValue string   `mapstructure:"PublishedValue" usage:"Value to which the filter property will be set after generating the content."`

	UseDateForFilename bool `mapstructure:"UseDateForFilename" usage:"Use the creation date to generate the post filename."`
	UseShortcodes      bool `mapstructure:"UseShortcodes" usage:"True if you want to generate shortcodes for unimplemented markdown blocks, such as callout or quote."`
}
