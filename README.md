# Umineko Bot

A Twitter bot that posts Umineko no Naku Koro ni character images with in-character opinions and quotes.

## Features

- **Random Mode**: Posts a random character image with an in-character opinion/quote
- **Erika Mode**: Posts only Erika images
- Day counter tracking
- Automated posting via GitHub Actions (09:00 UTC and 21:00 UTC)

## Setup

### 1. Clone and configure

```bash
git clone https://github.com/yourusername/umineko_bot.git
cd umineko_bot
cp .env.example .env
```

Edit `.env` with your Twitter API credentials:

```
TWITTER_API_KEY=your_api_key
TWITTER_API_SECRET=your_api_secret
TWITTER_ACCESS_TOKEN=your_access_token
TWITTER_ACCESS_TOKEN_SECRET=your_access_token_secret
```

### 2. Add images

Place character images in `data/images/` organized by character name:

```
data/images/
├── Beatrice/
├── battler/
├── erika/
├── bernkastel pc/
└── ...
```

### 3. Run locally

```bash
# Dry run (preview without posting)
go run . --dry-run --mode=random
go run . --dry-run --mode=erika

# Actually post
go run . --mode=random
go run . --mode=erika
```

## GitHub Actions

The bot runs automatically twice daily:

| Time (UTC) | Mode | Content |
|------------|------|---------|
| 09:00 | random | Character image + opinion |
| 21:00 | erika | Erika image only |

### Required Secrets

Add these in your repo's Settings > Secrets > Actions:

- `TWITTER_API_KEY`
- `TWITTER_API_SECRET`
- `TWITTER_ACCESS_TOKEN`
- `TWITTER_ACCESS_TOKEN_SECRET`

### Manual Trigger

You can manually trigger the workflow from the Actions tab and select which mode to run.

## Post Format

**Random mode:**
```
Day 42

Battler you dense motherfucker, SOLVE THE MYSTERY ALREADY

#Beatrice #UminekoNoNakuKoroNi
```

**Erika mode:**
```
[Just the image, no text]
```

## Adding Opinions

Edit `data/opinions.json` to add character opinions:

```json
{
  "character": "Beatrice",
  "opinions": [
    "Without love, it cannot be seen.",
    "Skill issue, Battler. Massive skill issue."
  ]
}
```

Character names must match the folder names in `data/images/` (case-insensitive).

## License

MIT
