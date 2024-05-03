# Video Encoding and Decoding Library

This repository contains a Go library for encoding and decoding video streams using the FFmpeg library, along with a client-server application that allows streaming video over a secure connection. The server encodes the incoming video stream using the H.264 codec, while the client decodes the received stream and renders it in a window.

## Features

- Convert raw image data to H.264 video frames
- Decode H.264 video frames into raw image data
- Support for multiple simultaneous connections
- Secure communication using TLS
- Screen capture and encoding on the client side
- Video streaming from the client to the server
- Video decoding and rendering on the client side

## Installation

1. Clone the repository:

```bash
git clone https://github.com/your-username/video-encoding-decoding.git
```

2. Install the required dependencies:

```bash
go install github.com/giorgisio/goav@latest
go install github.com/hashicorp/yamux@latest
go install github.com/faiface/pixel@latest
go install github.com/3d0c/gmfl@latest
go install github.com/kbinani/screenshot@latest
```

Вот исправленная версия README файла, включающая тесты для серверной и клиентской части:

## Setup

### Generating TLS Certificates

This application uses TLS for secure communication between the client and server. To generate a self-signed TLS certificate and key, run the following command:

```bash
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
```

This will create `cert.pem` (the certificate) and `key.pem` (the private key) files in the current directory. These files need to be placed on both the server and client machines, as they are required for establishing a secure TLS connection.

### Server Setup

1. Copy the `cert.pem` and `key.pem` files to the server machine.
2. Start the server:

```bash
go run main.go
```

The server will listen for incoming connections on `localhost:8000`.

### Client Setup

1. Copy the `cert.pem` and `key.pem` files to the client machine.
2. Run the client application:

```bash
go run main.go
```

The client will connect to the server at `133.133.133.133:8000` (replace with the appropriate server address).

## Usage

1. Start the server as described in the [Server Setup](#server-setup) section.
2. Run the client application as described in the [Client Setup](#client-setup) section.
3. The client will capture your screen, encode it into an H.264 video stream, and send it to the server.
4. The client will receive the encoded video stream from the server, decode it, and render it in a window.

## Testing

This project includes unit tests for various components of the library and the client application. The `tests.go` file contains the following tests:

### Server Tests

- `TestFrameToImage`: Tests the conversion of an AVFrame to an image.
- `TestHandleConnection`: Tests the handling of a client connection.
- `TestCreateVideoDecoder`: Tests the creation of a video decoder context.
- `TestEncodeVideo`: Tests the encoding of an image into a video frame.

### Client Tests

- `TestCaptureScreen`: Tests the screen capture functionality.
- `TestCreateVideoDecoder`: Tests the creation of a video decoder.
- `TestEncodeVideo`: Tests the encoding of a video frame.
- `TestFrameToImage`: Tests the conversion of a video frame to an image.
- `TestListenConnection`: Tests the listening for incoming video data.
- `TestRun`: Tests the main rendering loop.
- `TestSendVideo`: Tests the sending of a video packet.

To run the tests, execute the following command:

```bash
go test
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
