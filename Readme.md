# TARS - Your Copilot for Disk Space

<p align="center">
  <img src="Assets/icon.svg" width="96" height="96"
    alt="TARS icon - a Gargantua-style black hole with three concentric accretion disk rings wrapping around a central sphere, evoking the TARS robot from Interstellar. A directory tree is etched inside the sphere, representing the folder hierarchy TARS explores. The ring tracks mirror a disk platter, tying the space theme to the tool's purpose: understanding what lives on your storage."/>
</p>

> Understand what takes space, why it exists, what changed, and what your options mean.

**TARS** is a free, open-source macOS terminal app built with Go and Bubble Tea. It shows what occupies your disk, explains recognized files, tracks changes between scans, and supports deliberate cleanup - with clear consequences and no surprises. A later phase adds optional on-device AI explanations, no cloud account required.

*Inspired by the robot from Interstellar.*

---

## At a Glance

| | |
|---|---|
| **Command** | `tars` |
| **Platform** | macOS - Apple Silicon + Intel |
| **Stack** | Go · Bubble Tea · Lip Gloss · SQLite |
| **License** | [MIT](LICENSE) |
| **Min macOS** | macOS 13 Ventura+ |
| **Cost** | Free, forever |
| **Status** | Phase 0 complete — prototype next |

---

## Phases & Features

| Phase | What it delivers | Status |
|---|---|---|
| **0 · Prototype** | Interactive UI prototype with sample data; name/license decisions | `Done` |
| **1 · Explorer** | Real scan, disk usage, large-files view, Finder integration | `Planned` |
| **2 · Explanations** | App recognition, friendly item descriptions, read-only suggestions | `Planned` |
| **3 · History** | Opt-in SQLite scan history, growth tracking between scans | `Planned` |
| **4 · Cleanup** | Reviewed, confirmed cleanup with Trash support and result records | `Planned` |
| **5 · Release** | Signed/notarized builds, one-line launcher, accessibility, docs | `Planned` |
| **6 · Local AI** | On-device plain-language explanations - offline, no API key | `Planned` |
| **7 · Advanced** | Duplicate finder, goal-based cleanup, scheduled scans | `Candidate` |

---

## Key Features

- **Honest measurements** - separates logical size, allocated size, and system-reported usage
- **Item explanations** - friendly name, purpose, owning app, and confidence level
- **Change tracking** - see which folders grew since your last scan
- **Safe cleanup** - allowlist only; review → confirm → execute; no silent deletion
- **Local-first AI** - small on-device model, loaded on demand, fully offline
- **No setup required** - no Homebrew, Go, or Xcode needed to run

---

## Principles

- Explain before suggesting. Require explicit intent for cleanup.
- Show useful results without configuration.
- Keep scanning, recommendations, and history local and offline.
- Label what is measured, what is estimated, and what is unknown.
- Never delete silently. The UI cannot bypass the cleanup engine.

---

## Out of Scope

Windows/Linux · antivirus · general app uninstaller · health scores · automatic deletion · mandatory background services · mandatory AI or cloud

---

*Phase 0 (product definition) is complete. See [Docs/Decisions.md](Docs/Decisions.md) for all resolved decisions.*
