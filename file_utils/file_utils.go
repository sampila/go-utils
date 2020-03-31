package file_utils

import(
	"net/http"
  "mime/multipart"
)

func GetFileContentType(file *multipart.FileHeader) (string, error) {

  f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, readErr := f.Read(buffer)
	if readErr != nil {
		return "", err
	}
	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
