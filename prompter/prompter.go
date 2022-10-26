package prompter

import (
	"strings"

	"github.com/manifoldco/promptui"
)

type PromptSelection struct {
	DisplayName, Value string
}

type PromptSelectionOpts struct {
	Label string
	Size  int
}

func GetInputFromPrompt(selections []PromptSelection, opts *PromptSelectionOpts) (*PromptSelection, error) {
	prompt := promptui.Select{
		Label: opts.Label,
		Items: selections,
		Size:  opts.Size,
		Searcher: func(input string, index int) bool {
			selected := selections[index]
			name := strings.Replace(strings.ToLower(selected.DisplayName), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)

			return strings.Contains(name, input)
		},
		StartInSearchMode: true,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "\U0001F9A9 {{ .DisplayName | cyan }}",       // 選択したモノだけ色を変えたいなどの場合、選択したデータの出力フォーマット定義できます。
			Inactive: "  {{ .DisplayName | cyan }}",                // 選択されていない状態のフォーマット定義
			Selected: "\U0001F9A9 {{ .DisplayName | red | cyan }}", // 選択後のフォーマット定義
			Details: `
--------- device ----------
{{ "Name:" | faint }}   {{ .DisplayName }}
{{ "Value:" | faint }}  {{ .Value }}`,
		},
	}

	selectedIndex, _, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	return &selections[selectedIndex], nil
}
