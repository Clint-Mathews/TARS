# TARS — All Decisions Record

## Summary

| # | Decision | Answer |
|---|---|---|
| 1 | License | **MIT** |
| 2 | Minimum macOS | **macOS 13 Ventura** |
| 3 | Default scan scope | **Home folder (`~/`)** |
| 4 | Prototype tabs | **All visible, future ones "Coming Soon"** |
| 5 | Name (`tars`) | **Proceed** — available on Homebrew, no meaningful conflicts |
| 6 | First public release | **Read-only** (Phases 0–3 + 5, cleanup comes later) |
| 7 | Double-clickable launcher | **Not required** for first release; shell one-liner is enough |
| 8 | Release hosting | **GitHub Releases** |
| 9 | History retention | **Last 30 scans, capped at ~50MB** |
| 10 | History granularity | **Per-directory summaries** |
| 11 | AI runtime direction | **llama.cpp + ~1–3B quantized model** |
| 12 | Initial recognition rules | **14 categories** (see below) |
| 13 | Initial cleanup allowlist | **11 actions** (see below) |

---

## 1. License → MIT

- Matches entire dependency chain (Bubble Tea, Lip Gloss, Bubbles, llama.cpp)
- Simplest, most contributor-friendly
- **Action:** Commit a `LICENSE` file with MIT text + copyright line

---

## 2. Minimum macOS → macOS 13 Ventura (2022)

- Phase 1 is pinned to Go 1.27.1; Go 1.27 requires macOS 13+, exactly matching the selected support floor
- Supports Intel and Apple Silicon systems capable of running macOS 13; validate both architectures rather than claiming an unverified adoption percentage
- **Test matrix:**

| Axis | Values |
|---|---|
| macOS version | 13 Ventura, latest stable at validation time |
| Architecture | Apple Silicon (`arm64`), Intel (`amd64`) |
| Terminal | Terminal.app, iTerm2 |

---

## 3. Default Scan Scope → Home Folder (`~/`)

- No elevated permissions needed
- Matches the nontechnical user target audience
- Broader scans are explicit opt-in through the Phase 1 root override (for example, `tars --root /`)

---

## 4. Prototype Tabs → All Visible

- **Active in Phase 0:** Overview, Browse, Help
- **Shown as "Coming Soon":** Large Files, Apps & Data, Suggestions, History
- Gives usability testers a sense of the full product

---

## 5. Name Availability → Proceed with `tars`

| Registry | Status |
|---|---|
| Homebrew | ✅ Available as formula (no `tars` formula exists) |
| GitHub | ✅ `Clint-Mathews/TARS` is unique |
| pkg.go.dev | ✅ `github.com/Clint-Mathews/TARS` is unique |
| npm / crates.io / PyPI | ✅ No meaningful conflict (Go CLI) |
| Domain | ⏸️ Deprioritized — host on Vercel (free) if ever needed |

> ByteDance's Agent TARS / UI-TARS exist but are in a completely different product category (GUI automation AI vs. disk space explorer). No confusion risk for the target audience.

---

## 6. First Public Release → Read-Only

- Ships Phases 0–3 + 5 (prototype → explorer → explanations → history → distribution)
- **Phase 4 (cleanup) ships as a follow-up release**
- Rationale: lower risk, faster to ship, builds user trust before touching any files

---

## 7. Double-Clickable Launcher → Not Required

- First release uses a shell one-liner (`curl | sh` style with verification)
- Evaluate a `.app` wrapper later based on user feedback
- Keeps Phase 5 scope manageable

---

## 8. Release Hosting → GitHub Releases

- Free, already where the repo lives
- Supports signed artifacts and checksums
- Well-understood by the target audience
- Homebrew tap can be added later if demand warrants it

---

## 9. History Retention → 30 Scans / 50MB Cap

- **Default:** Keep the last 30 scan snapshots
- **Storage cap:** ~50MB max for the SQLite database
- Users can adjust both values through settings
- Rationale: 30 scans ≈ roughly monthly use for 2.5 years — generous without being wasteful

---

## 10. History Granularity → Per-Directory Summaries

- Each scan snapshot stores **per-directory**: total size, file count, classification
- Not per-file (too large) and not top-N only (too limited for drill-down)
- Good balance of growth-tracking usefulness vs. storage cost

---

## 11. AI Runtime Direction → llama.cpp + Small Quantized Model

- **Runtime:** llama.cpp (MIT licensed, proven Apple Silicon support via Metal)
- **Model range:** ~1–3B parameters quantized (e.g., Phi-3-mini, Gemma 2B)
- **Key constraints from requirements:**
  - Offline after setup, no cloud fallback
  - Load on demand, release resources after idle
  - Show download size, disk space, memory requirements before download
  - Model can explain but cannot authorize cleanup or bypass rules
- **Intel fallback:** CPU-only execution; keep the app fully usable if the model can't run comfortably
- Final model selection after hardware benchmarking in Phase 6

---

## 12. Initial Recognition Rules (Phase 2)

TARS will recognize and explain these categories from day one:

| # | Category | Path Pattern | Regenerable | Notes |
|---|---|---|---|---|
| 1 | Xcode Derived Data | `~/Library/Developer/Xcode/DerivedData` | ✅ | Rebuilds on next build |
| 2 | Homebrew Cache | `~/Library/Caches/Homebrew` | ✅ | Re-downloads on next install |
| 3 | npm / yarn / pnpm Cache | `~/.npm`, `~/Library/Caches/Yarn`, `~/.pnpm-store` | ✅ | Re-downloads on next install |
| 4 | CocoaPods Cache | `~/Library/Caches/CocoaPods` | ✅ | Re-downloads on next pod install |
| 5 | Gradle / Maven Cache | `~/.gradle/caches`, `~/.m2/repository` | ✅ | Re-downloads on next build |
| 6 | Docker Data | `~/Library/Containers/com.docker.docker` | ⚠️ | Images regenerable, volumes may not be |
| 7 | Spotify Cache | `~/Library/Caches/com.spotify.client` | ✅ | Re-streams content |
| 8 | Slack Cache | `~/Library/Application Support/Slack` | ✅ | Re-downloads from server |
| 9 | Browser Caches | `~/Library/Caches/com.google.Chrome`, `…/org.mozilla.firefox`, `…/com.apple.Safari` | ✅ | Pages reload from web |
| 10 | System Logs | `~/Library/Logs` | ⚠️ | Old logs generally safe; recent may be needed for debugging |
| 11 | macOS Software Updates | `/Library/Updates` | ✅ | Re-downloads if needed |
| 12 | iOS Device Backups | `~/Library/Application Support/MobileSync` | ❌ | Unique data, manual review only |
| 13 | Trash | `~/.Trash` | ❌ | Already user-deleted, but may still be wanted |
| 14 | Downloads Folder | `~/Downloads` | ❌ | Flag for review only, never auto-eligible |

**Additional:**
- Documents folder — recognized and explained but **always protected**
- User-specific / cloud repo-based items — recognize cloud-synced folders (iCloud, Dropbox, OneDrive) and flag them with sync status awareness

---

## 13. Initial Cleanup Allowlist (Phase 4)

When cleanup ships, these are the initial supported actions:

| # | Target | Action | Type |
|---|---|---|---|
| 1 | Xcode Derived Data | Trash entire folder | ✅ Auto-eligible |
| 2 | Homebrew Cache | Trash cache files | ✅ Auto-eligible |
| 3 | npm / yarn / pnpm Cache | Trash cache files | ✅ Auto-eligible |
| 4 | CocoaPods Cache | Trash cache files | ✅ Auto-eligible |
| 5 | Gradle / Maven Cache | Trash cache files | ✅ Auto-eligible |
| 6 | Spotify Cache | Trash cache folder | ✅ Auto-eligible |
| 7 | Slack Cache | Trash cache files | ✅ Auto-eligible |
| 8 | Browser Caches | Trash cache files | ✅ Auto-eligible |
| 9 | System Logs (>7 days) | Trash old log files | ✅ Auto-eligible |
| 10 | iOS Device Backups | Show size + date, user decides | ⚠️ Manual review |
| 11 | Docker Unused Images | Guide to `docker system prune` | ⚠️ Manual review |

**Additional:**
- Downloads folder — manual review (show age, size, let user pick individual files)
- Documents — manual review (user-driven, never bulk-eligible)
- User disk usage / cloud repos — informational only, direct to respective tools

**Always protected (never cleanup-eligible):** `~/Documents`, `~/Desktop`, `~/Photos`, `~/Music`, `~/.ssh`, `~/Library/Keychains`, `~/Library/Preferences`, iCloud Drive contents, any cloud-synced folder, system directories.

---

## Status: All Decisions Closed ✅

No open decisions remain. The project is unblocked for implementation across all phases.
