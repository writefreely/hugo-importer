# Hugo-to-WriteFreely Migration Tool

This command-line tool imports a content-source directory from a [Hugo] site
into a WriteFreely instance, including [Write.as].

## Summary

  - [Getting Started](#getting-started)
  - [Usage](#usage)
  - [Contributing](#contributing)
  - [Versioning](#versioning)
  - [Authors](#authors)

## Getting Started

These instructions will get you a copy of the project up and running on your
local machine for development and testing purposes. See deployment for notes
on how to deploy the project on a live system.

### Prerequisites

You'll need to have Go v1.15 or higher installed and available on your
system's `$PATH`. Please refer to the [Go installation instructions] for
detailed steps.

### Installing

First, start by cloning the project to your local machine:

```bash
git clone git@github.com:writefreely/hugo-importer.git
```

or, if you prefer cloning via HTTPS rather than SSH:

```bash
git clone https://github.com/writefreely/hugo-importer.git
```

Then, install the migration tool:

```bash
cd hugo-importer
go install
```

## Usage

Run the tool from the root directory of your Hugo site.

```bash
hugo-importer --user matt --blog my-great-blog --content-dir posts
```

The migration tool is interactive; it will prompt you for your password and to
confirm that you're ready to proceed, then output each step of the parsing and
publishing process as they occur.

Your password is not saved anywhere; it's used once to fetch an access token
that is used for subsequent requests to the WriteFreely/Write.as API.

Please review the [limitations](#limitations) below to confirm the migration
tool will work for your Hugo site.

### Parameters

Some parameters are required, and others only pertain to specific cases.

#### `--user` (or `-u`)

**Required**

The username for logging into your WriteFreely (or Write.as) account.

#### `--blog` (or `-b`)

**Required**

The alias of the WriteFreely/Write.as blog that the Hugo posts will be migrated
to; you can find this from the Blogs page on WriteFreely/Write.as, clicking on
the "Customize" button for that blog, and checking the **bold** portion in the
"URL" section ("Preferred URL" on Write.as).

#### `--content-dir` (or `-c`)

**Required**

The source directory for the posts you wish to migrate, under `content` in your
Hugo site directory. For example, to import `your-hugo-site/content/posts`, you
would set `--content-dir posts`.

#### `--instance` (or `-i`)

**Required for WriteFreely only**

The WriteFreely instance URL (e.g., `--instance https://pencil.writefree.ly)`;
leave this out if you're importing into a Write.as account.

#### `--images`

**Only available to Write.as Pro users with Snap.as**

By adding this flag, the migration tool will attempt to find local images when
parsing your source content, and upload them from your Hugo site's `static`
directory to [Snap.as].

### Commands

#### `help`

The above information is available via the `help` command:

```bash
hugo-importer help
```

### Output

Posts in the source directory are parsed, added to a queue, and then published
to the destination blog from the queue. The server's response for each post is
sanitized of the modification token and saved to a publishing log file in the
root directory of your Hugo site.

If the `--images` flag is set, and the migration tool encounters errors trying
to upload an image to Snap.as, it will note them in the console output and add
them to an `upload-error.log` file in the root directory of your Hugo site.

### Limitations

The migration tool expects a standard structure where `content` and `static`
directories are one level below the Hugo site's root directory, and that local
images are served out of the `static` directory.

The migration tool will convert _most_ [built-in Hugo shortcodes] into URLs, specifically:

- `gist`
- `instagram`
- `tweet`
- `vimeo`
- `youtube`

By converting these to URLs, `gist`, `tweet`, and `youtube` URLs are shown as
rich embeds for Write.as Pro users.

## Contributing

Please read the [CONTRIBUTING] documentation for details on how to report bugs,
submit feature requests, and for submitting pull requests.

## Versioning

We use [SemVer] for versioning. See the [tags on this repository] for versions
that are currently available.

## Authors

  - **[Angelo Stavrow]** - *Initial work and documentation*

See also the list of [contributors] who participated in this project.

<!--references-->
[Hugo]: https://gohugo.io/
[Write.as]: https://write.as/
[Go installation instructions]: https://golang.org/doc/install
[Snap.as]: https://snap.as
[built-in Hugo shortcodes]: https://gohugo.io/content-management/shortcodes/#use-hugos-built-in-shortcodes
[CONTRIBUTING]: CONTRIBUTING.md
[SemVer]: http://semver.org/
[tags on this repository]: https://github.com/writefreely/hugo-importer/tags
[Angelo Stavrow]: https://github.com/AngeloStavrow
[contributors]: https://github.com/WriteFreely/hugo-importer/contributors