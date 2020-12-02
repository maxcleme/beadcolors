# Bead Colors

## Context

Since several years, in the fusing beads community, we saw a bunch of Google Spreadsheet defining beads colors in color code formats. However these Sheets ofter lacks of updates or correctness.

This project aims to centralize, maintain and expose an exhaustive list of all beads colors references using GitHub.

## Support


| Brands | Color Count | Style |
| -------- | -------- | -------- | 
| Hama     | 56     | Melty Bead |
| Perler     | 78     | Melty Bead | 
| Perler Mini     |  41   | Melty Bead | 
| Perler Caps     |  26   | Bead | 
| Nabbi    | 29     | Melty Bead |
| Artkal A-2.6MM    | 144     | Melty Bead |
| Artkal C-2.6MM    | 144     | Melty Bead |
| Artkal R-5MM    | 88     | Melty Bead|
| Artkal S-5MM    | 183     | Melty Bead |
| Diamond Dotz    | 461     | Diamond Painting |


## Format

In order to prevent breaking change in the future, this repository is organized in several folders : 

| Folder | Description |
| -------- | -------- |
| `/raw`     | Very simple `.csv` for each brands, in the following format :<br/><br/> `[reference_code, name, rgb_r, rgb_g, rgb_b, contributor]` <br/><br/>**Note: Contribution should be made on this folder, orders folders will be generated with Github Actions.**|
| `/v1`     | `[reference_code, name, rgb_r, rgb_g, rgb_b, rgb_hex, contributor]` |
| `/v2`     | `[reference_code, name, rgb_r, rgb_g, rgb_b, hsl_h, hsl_s, hsl_l, lab_l, lab_a, lab_b, contributor]` |
| `/v3`     | `[reference_code, name, symbol, rgb_r, rgb_g, rgb_b, hsl_h, hsl_s, hsl_l, lab_l, lab_a, lab_b, contributor]` |



## Contribution

Simply create a PR applying modification under `/raw` folder **only**.

**If you are updating references, you should provide explanation about why your reference is more accurate than the current one.**


## Exposition


| Method |  Usage |
| -------- | -------- |
| Browsing website     | [Beadscolors](https://beadcolors.eremes.xyz/) |     
| Static files     | <ul><li>[Hama](https://beadcolors.eremes.xyz/raw/hama.csv)</li><li>[Nabbi](https://beadcolors.eremes.xyz/raw/nabbi.csv)</li><li>[Perler](https://beadcolors.eremes.xyz/raw/perler.csv)</li><li>[Perler Mini](https://beadcolors.eremes.xyz/raw/perler_mini.csv)</li><li>[Perler Caps](https://beadcolors.eremes.xyz/raw/perler_caps.csv)</li><li>[Artkal A-2.6MM](https://beadcolors.eremes.xyz/raw/artkal_a.csv)</li><li>[Artkal C-2.6MM](https://beadcolors.eremes.xyz/raw/artkal_c.csv)</li><li>[Artkal R-5MM](https://beadcolors.eremes.xyz/raw/artkal_r.csv)</li><li>[Artkal S-5MM](https://beadcolors.eremes.xyz/raw/artkal_s.csv)</li><li>[Diamond Dotz](https://beadcolors.eremes.xyz/raw/diamondDotz.csv)</li></ul>     |

## Automation


Goal of automation are multiple : 
1. Maintain only `raw` folder, contains minimal data, only `[ref, name, r, g, b, contributor]`
2. Enhance the `raw` folder by adding new fields to each color reference (eg: `hex`, `hsl`, `lab`)
3. Generate bunch of format (eg: `csv`, `json`, `yaml`) 

`v1` folder is more a proof of concept of what it could be done in a later time.
