package pleasanter

type AttachmentsHash map[string][]Attachment

type Attachment struct {
	Name        string `json:"Name"`
	ContentType string `json:"ContentType"`
	Base64      string `json:"Base64"`
	Deleted     int    `json:"Deleted"`
}

type CheckHash map[string]bool

type ClassHash map[string]string

type DateHash map[string]string

type DescriptionHash map[string]string

type NumHash map[string]float64
