# galfudim.dev

My personal website and blog built with [Hugo](https://gohugo.io/) and the [Ananke theme](https://github.com/theNewDynamic/gohugo-theme-ananke).

## 🚀 Quick Start

### Prerequisites
- [Hugo Extended](https://gohugo.io/installation/) (v0.100.0+)
- [Git](https://git-scm.com/)

### Local Development

1. Clone the repository:
   ```bash
   git clone https://github.com/galfudim/galfudim.dev.git
   cd galfudim.dev
   ```

2. Initialize and update submodules (for the theme):
   ```bash
   git submodule update --init --recursive
   ```

3. Start the Hugo development server:
   ```bash
   hugo server -D
   ```

4. Open your browser and visit `http://localhost:1313`

## 📁 Project Structure

```
├── archetypes/          # Content templates
├── assets/              # Assets (SCSS, JS, images)
├── content/             # Site content (Markdown files)
├── data/                # Data files
├── layouts/             # HTML templates
├── static/              # Static files
├── themes/              # Hugo themes
├── hugo.toml           # Hugo configuration
└── CNAME               # Custom domain configuration
```

## ✍️ Creating Content

### New Blog Post
```bash
hugo new content posts/my-new-post.md
```

### New Page
```bash
hugo new content pages/my-new-page.md
```

## 🏗️ Building for Production

```bash
hugo --minify
```

The built site will be in the `public/` directory.

## 🚀 Deployment

This site is automatically deployed to GitHub Pages when changes are pushed to the main branch.

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 🤝 Contributing

While this is a personal website, suggestions and improvements are welcome! Feel free to open an issue or submit a pull request.

---

Built with ❤️ using [Hugo](https://gohugo.io/)
