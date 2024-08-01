<img src="https://github.com/user-attachments/assets/13024f0e-f065-4b27-82f0-0803970ee954" width="600px"/>


# Go For It - TODO app
This is a go serverless rendered todo app project

The stack consists of:

**Backend:**
- Plain go backend `version > 1.22`, due to routing
- [a-h/templ](https://github.com/a-h/templ) for HTML template generation
- Gowebly for generating the stack + some helpers

**Frontend:**
- HTMX for reactivity
- TailwindCSS for pretty stuff

## Data store
Data store is currently a simple in-memory store, so its lost after restart of server

## Setup

> ❗️ Please make sure that you have installed the executable files for all the necessary tools before starting your project. Exactly:
>
> - `Air`: [https://github.com/air-verse/air](https://github.com/air-verse/air)

> - `Templ`: [https://github.com/a-h/templ](https://github.com/a-h/templ)
> - `golangci-lint`: [https://github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint)

To start your project, run the **Gowebly** CLI command in your terminal:

```console
gowebly run
```
