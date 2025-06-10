package controllers

import (
	"context"
	"e-com/modules/admin/req"
	"e-com/modules/admin/usecases"
	"e-com/pkg/utils"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	usecases usecases.ProductUsecase
}

func NewProductController(usecases usecases.ProductUsecase) *ProductController {
	return &ProductController{usecases: usecases}
}

func (cc *ProductController) Create(c *fiber.Ctx) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	priceStr := c.FormValue("price")
	categoryStr := c.FormValue("category")
	brandStr := c.FormValue("brand")
	sku := c.FormValue("sku")
	featuredStr := c.FormValue("featured")
	status := c.FormValue("status")
	stockStr := c.FormValue("stock")

	price, _ := strconv.Atoi(priceStr)
	category, _ := strconv.Atoi(categoryStr)
	brand, _ := strconv.Atoi(brandStr)
	stock, _ := strconv.Atoi(stockStr)
	featured, err := strconv.ParseBool(featuredStr)
	if err != nil {
		featured = false
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ไม่สามารถอ่านไฟล์ได้",
		})
	}
	files := form.File["images"]

	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)

	if err != nil {
		log.Fatal("Cloudinary init failed:", err)
	}

	var uploadedURLs []string
	ctx := context.Background()
	var wg sync.WaitGroup
	var mu sync.Mutex
	var uploadErr error

	for _, fileHeader := range files {
		wg.Add(1)
		go func(fh *multipart.FileHeader) {
			defer wg.Done()
			url, err := utils.UploadFileToCloudinary(ctx, cld, fileHeader)
			if err != nil {
				uploadErr = err
				return
			}
			mu.Lock()
			uploadedURLs = append(uploadedURLs, url)
			mu.Unlock()
		}(fileHeader)
	}

	wg.Wait()

	if uploadErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "อัพโหลดไฟล์ล้มเหลว",
			"error":   uploadErr.Error(),
		})
	}

	data := &req.ReqProduct{
		Name:        name,
		Description: description,
		Sku:         sku,
		Status:      status,
		Count:       stock,
		Price:       price,
		Category:    uint(category),
		Brand:       uint(brand),
		Images:      uploadedURLs,
		Featured:    featured,
	}

	product, err := cc.usecases.Create(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "เพิ่มสินค้าสำเร็จ",
		"data":    product,
	})
}

func (cc *ProductController) FindByQuery(c *fiber.Ctx) error {
	search := c.Query("search")
	filter := c.Query("sortBy", "newest")
	category, _ := strconv.Atoi(c.Query("category", "0"))
	brand, _ := strconv.Atoi(c.Query("brand", "0"))
	price, _ := strconv.Atoi(c.Query("priceRange", "0"))
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "12"))

	products, err := cc.usecases.FindByReq(filter, search, page, limit, category, brand, price)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    products,
	})
}

func (cc *ProductController) Update(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	name := c.FormValue("name")
	description := c.FormValue("description")
	priceStr := c.FormValue("price")
	categoryStr := c.FormValue("category")
	brandStr := c.FormValue("brand")
	sku := c.FormValue("sku")
	featuredStr := c.FormValue("featured")
	status := c.FormValue("status")
	stockStr := c.FormValue("stock")

	price, _ := strconv.Atoi(priceStr)
	category, _ := strconv.Atoi(categoryStr)
	brand, _ := strconv.Atoi(brandStr)
	stock, _ := strconv.Atoi(stockStr)
	featured, err := strconv.ParseBool(featuredStr)
	if err != nil {
		featured = false
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ไม่สามารถอ่านไฟล์ได้",
		})
	}
	files := form.File["images"]

	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)

	if err != nil {
		log.Fatal("Cloudinary init failed:", err)
	}

	var uploadedURLs []string
	ctx := context.Background()
	var wg sync.WaitGroup
	var mu sync.Mutex
	var uploadErr error

	for _, fileHeader := range files {
		wg.Add(1)
		go func(fh *multipart.FileHeader) {
			defer wg.Done()
			url, err := utils.UploadFileToCloudinary(ctx, cld, fileHeader)
			if err != nil {
				uploadErr = err
				return
			}
			mu.Lock()
			uploadedURLs = append(uploadedURLs, url)
			mu.Unlock()
		}(fileHeader)
	}

	wg.Wait()

	if uploadErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "อัพโหลดไฟล์ล้มเหลว",
			"error":   uploadErr.Error(),
		})
	}

	data := &req.ReqProduct{
		Name:        name,
		Description: description,
		Sku:         sku,
		Status:      status,
		Count:       stock,
		Price:       price,
		Category:    uint(category),
		Brand:       uint(brand),
		Images:      uploadedURLs,
		Featured:    featured,
	}

	updateProduct, err := cc.usecases.Update(uint(id), data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "อัพโหลดไฟล์ล้มเหลว",
			"error":   uploadErr.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"data":    updateProduct,
	})

}

func (cc *ProductController) FindOneById(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	data, err := cc.usecases.FindOneById(uint(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func extractPublicIDFromURL(url string) string {
	if idx := strings.Index(url, "?"); idx != -1 {
		url = url[:idx]
	}

	parts := strings.Split(url, "/")
	if len(parts) < 2 {
		return ""
	}

	fileName := parts[len(parts)-1]
	folder := parts[len(parts)-2]
	publicID := folder + "/" + strings.TrimSuffix(fileName, filepath.Ext(fileName))

	return publicID
}

func (cc *ProductController) DeleteImage(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	img, err := cc.usecases.FindImage(uint(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	if err != nil {
		panic(err)
	}

	publicID := extractPublicIDFromURL(img.Url)

	err = utils.DeleteImage(context.Background(), cld, publicID)
	if err != nil {
		log.Println("ลบรูปไม่สำเร็จ:", err)
	}

	if err := cc.usecases.DeleteImage(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    img,
	})
}

func (cc *ProductController) DeleteProduct(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	if err := cc.usecases.DeleteProduct(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}

func (cc *ProductController) FindTotal(c *fiber.Ctx) error {
	data, err := cc.usecases.FindTotal()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}
