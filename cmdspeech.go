package main

import (
	"context"
	"io"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var openaiApiKey string

func init() {
	speechCmd.Flags().StringVar(&openaiApiKey, "openai", "", "OpenAI API Key")
}

var speechCmd = &cobra.Command{
	Use:   "speech",
	Short: "speech",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := openai.NewClient(openaiApiKey)

		request := openai.CreateSpeechRequest{
			Model: openai.TTSModel1,
			Input: "パリはフランスの首都であり、美しい建築と豊かな文化で世界中から観光客を魅了しています。代表的な観光地には、エッフェル塔があります。高さ約330メートルのこの鉄塔は、パリの象徴として知られています。また、世界最大級の美術館であるルーヴル美術館では、「モナ・リザ」や「ミロのヴィーナス」などの名画を鑑賞できます。ノートルダム大聖堂はゴシック建築の傑作として有名です。さらに、シャンゼリゼ通りと凱旋門はパリの華やかさを象徴しています。セーヌ川沿いの景色も魅力的で、クルーズ船からパリの美しい街並みを楽しめます。パリは歴史と芸術、グルメが融合する魅力的な都市です",
			Voice: openai.VoiceEcho,
			Speed: 1.7,
			ResponseFormat: openai.SpeechResponseFormatMp3,
		}
		res, err := client.CreateSpeech(context.Background(), request)
		if err != nil {
			return err
		}
		f, err := os.Create("a.mp3")
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := io.Copy(f, res); err != nil {
			return err
		}
		return nil
	},
}
