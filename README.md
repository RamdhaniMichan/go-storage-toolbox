# Go Storage Toolbox
Welcome to the Go Storage Toolbox, a versatile Go package that simplifies file storage management across various cloud and local storage providers. Whether you're building a web application, data pipeline, or a personal project, this toolbox has you covered.

## Features

- **Multi-Provider Support:** Seamlessly switch between different storage providers:
    - [Google Cloud Storage (GCS)](https://cloud.google.com/storage)
    - [MinIO](https://min.io)
    - Add your own providers with ease!

- **Upload & Download:** Easily upload and download files to/from your chosen storage provider with just a few lines of code.

- **Flexible Configuration:** Configure your storage providers using environment variables or configuration files.

## Usage

```

// Initsialize for Provider
// Example for Download File with MinIO)
err := NewStorageProvider().InitMiniO().DownloadFile({{bucket-name}}, {{path/filename}}, {{out-file}})
if err != nil {
  log.Fatal(err)
  return
}

```