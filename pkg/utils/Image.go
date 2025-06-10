package utils

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadFileToCloudinary(ctx context.Context, cld *cloudinary.Cloudinary, fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	uploadParams := uploader.UploadParams{
		Folder: "e-com",
	}

	uploadResult, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}

func DeleteImage(ctx context.Context, cld *cloudinary.Cloudinary, publicID string) error {
	_, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	return err
}
