=====

Bugs & Shortcomings
-------------------

- Support UNH, UNT etc.; remove warning in segspec.html once supported
- bug? language name code (3453) missing from UNCL spec.   

Features
--------

- Tests & benchmarks
  - Add github.com/PuerkitoBio/goquery library for validating response body
- Sort code search results
- Filtering of spec tables?
- Expose additional spec info (notes, full description)
- File upload of spec
- REST API for validating messages?
- edify: give list of available (downloaded/extracted) spec versions
- Spec search: fill combo box with available versions (default: most recent)
- structured output of parsed message (Dynatree?)
- Render parsed nested msg with segment links
- Display of (original?) message with tooltips containing spec information?
- Rendering of messages with template
  - future (maybe): template as form, to allow creation of message
- store uploaded messages in DB?
- Ansible scripts for public server
- Parallel search for specs?
  - Not pressing; search with all results takes 32ms, with no results: 1.4 ms
  - possibly: parallel rendering of templates?
  - alternative: render template in response to search; this template
    contains multiple tabs; each tab loads search results for a particular
    spec type.
- Spec search: support codes
- Spec search: organize different kinds of specs in accordion or tabs
  - display badge with result count
  - omit spec types with no results?
- Spec search. (maybe) limit search to particular kinds of specs
- Pagination of long tables (example: long code list http://localhost:8001/specs/simpledataelement/4451)
  - client-side pagination for simplicity
  - link at bottom of page to jump to top of page
- Autocomplete for search?



