package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/arjendevos/elevenlabs/client/types"
)

func (c Client) TTSWriter(ctx context.Context, w io.Writer, text, modelID, voiceID string, options types.SynthesisOptions) error {
	options.Clamp()
	url := fmt.Sprintf(c.endpoint+"/v1/text-to-speech/%s", voiceID)
	opts := types.TTS{
		Text:          text,
		ModelID:       modelID,
		VoiceSettings: options,
	}
	b, _ := json.Marshal(opts)
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("xi-api-key", c.apiKey)
	req.Header.Set("User-Agent", "github.com/arjendevos/elevenlabs")
	req.Header.Set("accept", "audio/mpeg")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	switch res.StatusCode {
	case 401:
		return ErrUnauthorized
	case 200:
		defer res.Body.Close()
		io.Copy(w, res.Body)
		return nil
	case 422:
		fallthrough
	default:
		ve := types.ValidationError{}
		defer res.Body.Close()
		jerr := json.NewDecoder(res.Body).Decode(&ve)
		if jerr != nil {
			err = ErrorsJoin(err, jerr)
		} else {
			err = ErrorsJoin(err, ve)
		}
		return err
	}
}

func (c Client) TTS(ctx context.Context, text, voiceID, modelID string, options types.SynthesisOptions) ([]byte, error) {
	options.Clamp()
	url := fmt.Sprintf(c.endpoint+"/v1/text-to-speech/%s", voiceID)
	client := &http.Client{}
	opts := types.TTS{
		Text:          text,
		ModelID:       modelID,
		VoiceSettings: options,
	}
	b, _ := json.Marshal(opts)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("xi-api-key", c.apiKey)
	req.Header.Set("User-Agent", "github.com/arjendevos/elevenlabs")
	req.Header.Set("accept", "audio/mpeg")
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	fmt.Println("res", res.StatusCode)
	switch res.StatusCode {
	case 401:
		return []byte{}, ErrUnauthorized
	case 200:
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("Failed to read response body: %v", err)
		}

		defer res.Body.Close()
		return data, nil
	case 422:
		fallthrough
	default:
		ve := types.ValidationError{}
		defer res.Body.Close()
		jerr := json.NewDecoder(res.Body).Decode(&ve)
		if jerr != nil {
			err = ErrorsJoin(err, jerr)
		} else {
			err = ErrorsJoin(err, ve)
		}
		return []byte{}, err
	}
}

func (c Client) TTSStream(ctx context.Context, w io.Writer, text, voiceID string, options types.SynthesisOptions) error {
	options.Clamp()
	url := fmt.Sprintf(c.endpoint+"/v1/text-to-speech/%s/stream", voiceID)
	opts := types.TTS{
		Text:          text,
		VoiceSettings: options,
	}
	b, _ := json.Marshal(opts)
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("xi-api-key", c.apiKey)
	req.Header.Set("User-Agent", "github.com/arjendevos/elevenlabs")
	req.Header.Set("accept", "audio/mpeg")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	switch res.StatusCode {
	case 401:
		return ErrUnauthorized
	case 200:
		defer res.Body.Close()
		io.Copy(w, res.Body)
		return nil
	case 422:
		fallthrough
	default:
		ve := types.ValidationError{}
		defer res.Body.Close()
		jerr := json.NewDecoder(res.Body).Decode(&ve)
		if jerr != nil {
			err = ErrorsJoin(err, jerr)
		} else {
			err = ErrorsJoin(err, ve)
		}
		return err
	}
}
