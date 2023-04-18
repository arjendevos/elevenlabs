package client

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/taigrr/elevenlabs/client/types"
)

func (c Client) DeleteVoiceSample(ctx context.Context, voiceID, sampleID string) (bool, error) {
	url := fmt.Sprintf(c.endpoint+"/v1/voices/%s/samples/%s", voiceID, sampleID)
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("xi-api-key", c.apiKey)
	req.Header.Set("User-Agent", "github.com/taigrr/elevenlabs")
	res, err := client.Do(req)

	switch res.StatusCode {
	case 422:
		ve := types.ValidationError{}
		defer res.Body.Close()
		jerr := json.NewDecoder(res.Body).Decode(&ve)
		if jerr != nil {
			err = errors.Join(err, jerr)
		} else {
			err = errors.Join(err, ve)
		}
		return false, err
	case 401:
		return false, ErrUnauthorized
	case 200:
		if err != nil {
			return false, err
		}
		return true, nil
	default:
		return false, errors.Join(err, ErrUnspecified)
	}
}

func (c Client) DownloadVoiceSampleWriter(ctx context.Context, w io.Writer, voiceID, sampleID string) error {
	url := fmt.Sprintf(c.endpoint+"/v1/voices/%s/samples/%s/audio", voiceID, sampleID)
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("xi-api-key", c.apiKey)
	req.Header.Set("User-Agent", "github.com/taigrr/elevenlabs")
	req.Header.Set("accept", "audio/mpeg")
	res, err := client.Do(req)

	switch res.StatusCode {
	case 422:
		ve := types.ValidationError{}
		defer res.Body.Close()
		jerr := json.NewDecoder(res.Body).Decode(&ve)
		if jerr != nil {
			err = errors.Join(err, jerr)
		} else {
			err = errors.Join(err, ve)
		}
		return err
	case 401:
		return ErrUnauthorized
	case 200:
		if err != nil {
			return err
		}
		defer res.Body.Close()
		io.Copy(w, res.Body)
		return nil
	default:
		return errors.Join(err, ErrUnspecified)
	}
}

func (c Client) DownloadVoiceSample(ctx context.Context, voiceID, sampleID string) ([]byte, error) {
	url := fmt.Sprintf(c.endpoint+"/v1/voices/%s/samples/%s/audio", voiceID, sampleID)
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("xi-api-key", c.apiKey)
	req.Header.Set("User-Agent", "github.com/taigrr/elevenlabs")
	req.Header.Set("accept", "audio/mpeg")
	res, err := client.Do(req)

	switch res.StatusCode {
	case 422:
		ve := types.ValidationError{}
		defer res.Body.Close()
		jerr := json.NewDecoder(res.Body).Decode(&ve)
		if jerr != nil {
			err = errors.Join(err, jerr)
		} else {
			err = errors.Join(err, ve)
		}
		return []byte{}, err
	case 401:
		return []byte{}, ErrUnauthorized
	case 200:
		if err != nil {
			return []byte{}, err
		}
		b := bytes.Buffer{}
		w := bufio.NewWriter(&b)

		defer res.Body.Close()
		io.Copy(w, res.Body)
		return b.Bytes(), nil
	default:
		return []byte{}, errors.Join(err, ErrUnspecified)
	}
}