package utils

var (
	modelShorteners map[string]string = map[string]string{
		"phi-2": "github://mudler/LocalAI/examples/configurations/phi-2.yaml@master",
		"llava": "github://mudler/LocalAI/examples/configurations/llava.yaml@master",
	}
)

func ModelShortURL(s string) string {
	if _, ok := modelShorteners[s]; ok {
		s = modelShorteners[s]
	}

	return s
}
