# KAP Data Publishing Service REST API

KAP (Public Disclosure Platform) REST APIs provide secure, authorized access to
KAP data. The API follows a REST architecture and returns JSON responses.

## Table of Contents

- [Overview](#overview)
- [Authentication](#authentication)
- [Base URLs](#base-urls)
- [Error Handling](#error-handling)
- [Disclosure Class and Type Reference](#disclosure-class-and-type-reference)
- [Endpoints](#endpoints)
  - [Generate Token](#generate-token)
  - [Disclosure List](#disclosure-list)
  - [Disclosure Detail](#disclosure-detail)
  - [Download Attachment](#download-attachment)
  - [Last Disclosure Index](#last-disclosure-index)
  - [Corporate Action Event Status](#corporate-action-event-status)
  - [Blocked Disclosures](#blocked-disclosures)
  - [Company List](#company-list)
  - [Company Securities](#company-securities)
  - [Company Detail](#company-detail)
  - [Fund List](#fund-list)
  - [Fund Detail](#fund-detail)
- [Data Notes](#data-notes)

## Overview

| Property    | Value                                      |
|-------------|--------------------------------------------|
| API Title   | KAP VYK API                                |
| Version     | 0.0.1                                      |
| Format      | REST / JSON                                |
| Auth        | API Key + Bearer Token (production)        |
| Auth (test) | Basic Authentication (no token required)   |

## Authentication

### Production Environment

All production requests require a bearer token obtained from the
[Generate Token](#generate-token) endpoint.

1. Call `/auth/generateToken` with your `apiKey` to receive a token.
2. Include the token in the `Authorization` header for all subsequent requests.
3. Each token is valid for **24 hours**. After expiry, generate a new one.
4. Multiple tokens may be active simultaneously.

**Header format:**

```text
Authorization: <token>
Content-Type: application/json
```

### Test Environment

No token is required for the test environment. Basic authentication is used
with read/write scopes.

### Invalid Token Response

If a request is made with an expired or invalid token:

```json
{
  "code": "ER006",
  "message": "Token geçerlilik süresi bitmiştir. The token has expired. Please try again with a valid token."
}
```

## Base URLs

| Environment | URL                            |
|-------------|--------------------------------|
| Test        | `https://apigwdev.mkk.com.tr`  |

All endpoint paths are relative to the base URL.

## Error Handling

All error responses (HTTP 400 and 500) return the following structure:

```json
{
  "code": "string",
  "message": "string"
}
```

### Error Codes

| Code  | Message                                                      | Description                                                                                      |
|-------|--------------------------------------------------------------|--------------------------------------------------------------------------------------------------|
| ER001 | Servis erişim yetkiniz bulunmamaktadır.                      | The data distribution firm has no access permission to any service.                              |
| ER002 | Unauthorized request.                                        | The data distribution firm has no access permission to the customized service.                   |
| ER003 | Servis erişim yetkiniz bulunmamaktadır.                      | An access attempt was made from an unregistered IP address.                                      |
| ER004 | Unauthorized request.                                        | The token is invalid.                                                                            |
| ER005 | Ip bilgisi doğrulanamadı.                                    | No disclosure was found in the system with the queried ID.                                       |
| ER006 | Sender ip verification failed for the sender ip.             | Invalid token warning.                                                                           |
| ER007 | Token bilgisi doğrulanamadı.                                 | IP information could not be found.                                                               |
| ER008 | Authorization token is not valid.                            | No disclosure was found in the system with the queried ID (for customized data publishing service access). |

## Disclosure Class and Type Reference

### Disclosure Class (`disclosureClass`)

| Code | Turkish                            | English                              |
|------|------------------------------------|--------------------------------------|
| FR   | Finansal Rapor Bildirimi           | Financial Report Disclosure          |
| ODA  | Özel Durum Açıklaması Bildirimi    | Material Event Disclosure            |
| DG   | Diğer Bildirim                     | Other Disclosure                     |
| DUY  | Düzenleyici Kurum Bildirimi        | Regulatory Authority Disclosure      |

### Disclosure Type (`disclosureType`)

| Code | Turkish                            | English                              |
|------|------------------------------------|--------------------------------------|
| FR   | Finansal Rapor Bildirimi           | Financial Report Disclosure          |
| ODA  | Özel Durum Açıklaması Bildirimi    | Material Event Disclosure            |
| DG   | Diğer Bildirim                     | Other Disclosure                     |
| DUY  | Düzenleyici Kurum Bildirimi        | Regulatory Authority Disclosure      |
| FON  | Fon Bildirimi                      | Fund Disclosure                      |
| CA   | Hak Kullanım Bildirimi             | Corporate Action Disclosure          |

## Endpoints

### Generate Token

Generates a bearer token for production API access.

| Property | Value                                                                  |
|----------|------------------------------------------------------------------------|
| URL      | `/auth/generateToken`                                                  |
| Method   | `GET`                                                                  |
| Version  | 1.0                                                                    |
| Host     | `https://apigwdev.mkk.com.tr`                                         |

> This endpoint is only required for the production environment. The test
> environment does not require token generation.

#### Request Parameters

| Field    | Type   | Description | Required | Format       |
|----------|--------|-------------|----------|--------------|
| `apiKey` | String | API key     | Yes      | `[A-Z0-9]{36}` |

#### Example Request

```text
GET /auth/generateToken?apiKey=29223dec-32bc-49fb-919f-51405d110ab2
```

#### Response Parameters

| Field   | Type   | Description                 | Required | Format    |
|---------|--------|-----------------------------|----------|-----------|
| `token` | String | System-generated JWT token  | Yes      | `[A-Z0-9]` |

#### Example Response

```json
{
  "token": "eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJBWUwiLCJpYXQiOjE2ODkwNzU2ODUsImV4cCI6MTY4OTA3OTI4NX0.1DDmdKP74vPgW40XYMYiwbX73g9M7MbgsxWft_EJ3Rle4aqeHmoXeMvpRDGZyBPTqwdkbM1AFFUVv0CMjVgrGg"
}
```

---

### Disclosure List

Returns the first 50 disclosures starting from the given `disclosureIndex`,
including their type, class, disclosure ID, and report ID.

| Property | Value                                                                  |
|----------|------------------------------------------------------------------------|
| URL      | `/api/vyk/disclosures`                                                 |
| Method   | `GET`                                                                  |
| Version  | 1.0                                                                    |
| Host     | `https://apigwdev.mkk.com.tr`                                         |

#### Request Parameters

| Field              | Type    | Description                                                                                          | Required | Format              |
|--------------------|---------|------------------------------------------------------------------------------------------------------|----------|---------------------|
| `disclosureIndex`  | Integer | Starting disclosure index for listing. Starts from 538004. Disclosures 84196-538004 are pre-KAP 4.0. | Yes      | `[0-9]{max 10}`     |
| `disclosureClass`  | String  | Disclosure class filter (see [reference](#disclosure-class-disclosureclass))                          | No       | `[A-Z0-9]{max 10}`  |
| `disclosureType`   | String  | Disclosure type filter (see [reference](#disclosure-type-disclosuretype))                             | No       | `[A-Z0-9]{max 10}`  |
| `companyId`        | String  | Company ID filter                                                                                    | No       | `[0-9]`             |

#### Example Request

```text
GET /api/vyk/disclosures?disclosureIndex=1092228&disclosureTypes=DG&disclosureClass=DG&companyId=4329
```

#### Response Parameters

| Field                   | Type   | Description                                                                | Required | Format              |
|-------------------------|--------|----------------------------------------------------------------------------|----------|---------------------|
| `disclosureIndex`       | String | Disclosure publication number, starts from the requested index             | Yes      | `[0-9]{max 10}`     |
| `disclosureType`        | String | Disclosure type                                                            | Yes      | `[A-Z0-9]{max 10}`  |
| `disclosureClass`       | String | Disclosure class                                                           | Yes      | `[A-Z0-9]{max 10}`  |
| `subReportIds`          | List   | Sub-reports of the disclosure. Multiple entries possible. Sent for disclosures with `presentation` data file type. | Yes | `[A-Z0-9]` |
| `title`                 | String | Disclosure sender company title                                            | Yes      | `[A-Z0-9]`          |
| `companyId`             | Integer| Disclosure sender company ID                                               | Yes      | `[0-9]`             |
| `fundId`                | Integer| Fund ID (only present for fund disclosures)                                | Yes      | `[0-9]`             |
| `fundCode`              | String | Fund code (only present for fund disclosures)                              | Yes      | `[A-Z]`             |
| `acceptedDataFileTypes` | List   | Available file types for the disclosure: `html`, `presentation`            | Yes      | `[A-Z]`             |

#### Example Response

```json
[
  {
    "disclosureIndex": "1092228",
    "disclosureType": "DG",
    "disclosureClass": "DG",
    "subReportIds": [
      "oda-22300_Prospectus-Summary"
    ],
    "title": "VAKIF VARLIK KİRALAMA A.Ş.",
    "companyId": "4329",
    "acceptedDataFileTypes": [
      "html",
      "presentation"
    ]
  }
]
```

---

### Disclosure Detail

Returns the details of a disclosure for the given disclosure index.

| Property | Value                                                                  |
|----------|------------------------------------------------------------------------|
| URL      | `/api/vyk/disclosureDetail/{disclosureIndex}`                          |
| Method   | `GET`                                                                  |
| Version  | 1.0                                                                    |
| Host     | `https://apigwdev.mkk.com.tr`                                         |

Disclosures may contain sub-reports. To retrieve specific sub-report details,
pass report IDs in the `subReportList` parameter alongside the
`disclosureIndex`. If `subReportList` is not provided, all sub-reports are
returned.

Disclosures with index 84196-538004 (pre-KAP 4.0) can only be retrieved as
`html`. Index 538004 and above supports both `html` and `data` file types.

#### Request Parameters

| Field              | Type    | Description                                                                          | Required | Format          |
|--------------------|---------|--------------------------------------------------------------------------------------|----------|-----------------|
| `disclosureIndex`  | Integer | Disclosure index to query (path parameter)                                           | Yes      | `[0-9]{max 10}` |
| `fileType`         | String  | Desired file type: `html` or `data`                                                  | Yes      | `[A-Z]{4}`      |
| `subReportList`    | String  | Specific sub-report ID(s) to retrieve. Omit to get all sub-reports.                  | No       | `[A-Z0-9]`      |

**Test data note:** Disclosures with index between 1091689 and 1231017 are
available for querying in the test environment.

#### Example Request

```text
GET /api/vyk/disclosureDetail/1211180?fileType=data
```

#### Response Parameters

| Field                      | Type   | Description                                                                                  | Required |
|----------------------------|--------|----------------------------------------------------------------------------------------------|----------|
| `disclosureIndex`          | Integer| Disclosure publication number                                                                | Yes      |
| `senderId`                 | String | Publishing company member ID                                                                 | Yes      |
| `senderTitle`              | String | Publishing company member title                                                              | Yes      |
| `senderExchCodes`          | List   | Publishing company stock codes                                                               | Yes      |
| `behalfSenderId`           | String | Member ID on whose behalf the disclosure is published                                        | No       |
| `behalfSenderTitle`        | String | Member title on whose behalf the disclosure is published                                     | No       |
| `behalfSenderExchCodes`    | List   | Stock codes of the company on whose behalf the disclosure is published                       | No       |
| `behalfFundCode`           | String | Fund code on whose behalf the disclosure is published                                        | No       |
| `behalfFundTitle`          | String | Fund title on whose behalf the disclosure is published                                       | No       |
| `disclosureReason`         | String | Publication reason: `NEW` (New), `UPD` (Update), `CORR` (Correction), `CANC` (Cancellation) | Yes      |
| `disclosureDelayStatus`    | String | Delayed disclosure indicator: `O` (Not delayed), `D` (Delayed)                              | No       |
| `relatedDisclosureIndex`   | String | Related disclosure number (for updates, corrections, or cancellations)                       | No       |
| `disclosureType`           | String | Disclosure type                                                                              | Yes      |
| `disclosureClass`          | String | Disclosure class                                                                             | Yes      |
| `subject`                  | Object | Disclosure subject with `tr` (Turkish) and `en` (English) fields                             | Yes      |
| `consolidation`            | String | Consolidation status (FR class only): `CS` (Consolidated), `NC` (Unconsolidated)             | No       |
| `year`                     | Integer| Year the disclosure belongs to (periodic disclosures)                                        | No       |
| `period`                   | Object | Period with `tr` and `en` fields (e.g., "3 Months", "6 Months", "9 Months", "Annual")       | No       |
| `relatedStocks`            | List   | Related companies/funds mentioned in the disclosure, each with a `code` field                | No       |
| `summary`                  | Object | Disclosure summary with `tr` and `en` fields                                                 | Yes      |
| `time`                     | String | Publication time (format: `dd.MM.yyyy HH:mm:ss`)                                            | Yes      |
| `link`                     | String | Publication link on KAP                                                                      | Yes      |
| `attachmentUrls`           | List   | Disclosure attachments, each with `url` and `fileName`. Only PDF files.                      | Yes      |
| `eventType`                | String | Corporate action event type (only for CA disclosures)                                        | No       |
| `eventId`                  | Integer| Corporate action event ID (same for all disclosures in a process until completion)           | No       |
| `presentation`             | List   | Structured content with `id` and `content` fields                                            | Yes      |
| `flatData`                 | List   | Flat data content with `id` and `content` fields                                             | Yes      |
| `htmlMessages`             | List   | HTML message content with `id`, `tr`, `en` fields. Content is base64-encoded.                | Yes      |

#### Example Response

```json
{
  "disclosureIndex": "1211180",
  "senderId": "926",
  "senderTitle": "ECZACIBAŞI YATIRIM HOLDİNG ORTAKLIĞI A.Ş.",
  "senderExchCodes": ["ECZYT"],
  "behalfSenderId": "926",
  "behalfSenderTitle": "ECZACIBAŞI YATIRIM HOLDİNG ORTAKLIĞI A.Ş.",
  "behalfSenderExchCodes": ["ECZYT"],
  "disclosureReason": "CORR",
  "relatedDisclosureIndex": "1211162",
  "disclosureType": "FR",
  "disclosureClass": "FR",
  "subject": {
    "tr": "Faaliyet Raporu (Konsolide Olmayan)",
    "en": "Operating Review (Unconsolidated)"
  },
  "year": "2023",
  "period": {
    "tr": "9 Aylık",
    "en": "9 Months"
  },
  "relatedStocks": [],
  "summary": {
    "tr": "Faaliyet Raporu",
    "en": null
  },
  "time": "29.10.2023 14:05:18",
  "link": "https://kapsitealpha.mkk.com.tr/Bildirim/1211180",
  "attachmentUrls": [
    {
      "url": "https://vykapialpha.mkk.com.tr/api/vyk/downloadAttachment/4028328d8b2fcee7018b7aea7e3c631f",
      "fileName": "EYH Faaliyet Raporu 30.09.2023 .pdf"
    }
  ],
  "presentation": [
    {
      "id": "oda-34000_Unconsolidated-Operating-Review",
      "content": {
        "id": "oda-34000_Unconsolidated-Operating-Review",
        "isMultiDimensional": "yes",
        "ContextList": {
          "Context": {
            "id": "2023-10-29",
            "key": "CURR",
            "Period": {
              "instant": "2023-10-29"
            }
          }
        }
      }
    }
  ]
}
```

---

### Download Attachment

Downloads disclosure attachments using the URL provided in the
`attachmentUrls` field of the [Disclosure Detail](#disclosure-detail) response.

| Property | Value                                                                  |
|----------|------------------------------------------------------------------------|
| URL      | `/api/vyk/downloadAttachment/{id}`                                     |
| Method   | `GET`                                                                  |
| Version  | 1.0                                                                    |
| Host     | `https://apigwdev.mkk.com.tr`                                         |

> Attachment URLs from the disclosure detail response are not directly
> accessible to external users. Download attachments to your own storage and
> serve them from your own system.

#### Request Parameters

| Field | Type   | Description                                                                              | Required |
|-------|--------|------------------------------------------------------------------------------------------|----------|
| `id`  | String | Attachment ID from the `attachmentUrls[].url` field of the disclosure detail response    | Yes      |

#### Example Request

```text
GET /api/vyk/downloadAttachment/4028328d8b2fcee7018b7aea7e3c631f
```

#### Response

Binary file content. The attachment is rendered directly.

---

### Last Disclosure Index

Returns the ID of the most recently published disclosure.

| Property | Value                                                                  |
|----------|------------------------------------------------------------------------|
| URL      | `/api/vyk/lastDisclosureIndex`                                         |
| Method   | `GET`                                                                  |
| Version  | 1.0                                                                    |
| Host     | `https://apigwdev.mkk.com.tr`                                         |

#### Request Parameters

None.

#### Example Request

```text
GET /api/vyk/lastDisclosureIndex
```

#### Response Parameters

| Field                  | Type    | Description                            | Required | Format          |
|------------------------|---------|----------------------------------------|----------|-----------------|
| `lastDisclosureIndex`  | Integer | ID of the most recently published disclosure | Yes | `[0-9]{max 10}` |

#### Example Response

```json
{
  "lastDisclosureIndex": "1231017"
}
```

---

### Corporate Action Event Status

Returns the status of queried corporate action processes.

| Property | Value                                                                  |
|----------|------------------------------------------------------------------------|
| URL      | `/api/vyk/caEventStatus`                                               |
| Method   | `GET`                                                                  |
| Version  | 1.0                                                                    |
| Host     | `https://apigwdev.mkk.com.tr`                                         |

#### Request Parameters

| Field          | Type   | Description                       | Required |
|----------------|--------|-----------------------------------|----------|
| `processRefId` | String | Corporate action process reference ID | Yes  |

#### Response Parameters

| Field          | Type   | Description                           | Required |
|----------------|--------|---------------------------------------|----------|
| `refId`        | String | Process reference ID                  | Yes      |
| `status`       | String | Process status                        | Yes      |
| `statusReason` | String | Reason for the status                 | No       |
| `completeDate` | String | Process completion date               | No       |

---

### Blocked Disclosures

Returns the list of disclosures and disclosure attachments that have been
blocked (access restricted) in the KAP system up to the current date.

| Property | Value                                                                  |
|----------|------------------------------------------------------------------------|
| URL      | `/api/vyk/blockedDisclosures`                                          |
| Method   | `GET`                                                                  |
| Version  | 1.0                                                                    |
| Host     | `https://apigwdev.mkk.com.tr`                                         |

#### Request Parameters

None.

#### Response

Returns a `BlockedBase` object containing the list of blocked disclosures.

---

### Company List

Returns the list of all KAP member companies.

| Property | Value                                                                  |
|----------|------------------------------------------------------------------------|
| URL      | `/api/vyk/members`                                                     |
| Method   | `GET`                                                                  |
| Version  | 1.0                                                                    |
| Host     | `https://apigwdev.mkk.com.tr`                                         |

#### Request Parameters

None.

#### Response Parameters

| Field        | Type    | Description                             | Required | Format              |
|--------------|---------|-----------------------------------------|----------|---------------------|
| `id`         | Integer | Unique company ID                       | Yes      | `[0-9]{max 10}`     |
| `title`      | String  | Company title                           | Yes      | `[A-Z0-9]`          |
| `stockCode`  | String  | Company stock code(s). Multiple possible. | Yes    | `[A-Z0-9]`          |
| `memberType` | String  | Company type (see table below)          | Yes      | `[A-Z0-9]{max 10}`  |
| `kfifUrl`    | String  | Participation finance principles info page URL. Present for listed companies with active stock trading. | No | `[A-Z0-9]` |

**Member Types:**

| Code | Turkish                           | English                              |
|------|-----------------------------------|--------------------------------------|
| IGS  | İşlem Gören Şirket                | Listed Company                       |
| IGMS | İşlem Görmeyen Şirket             | Unlisted Company                     |
| YK   | Yatırım Kuruluşu                  | Investment Firm                      |
| PYS  | Portföy Yönetim Şirketi           | Portfolio Management Company         |
| DDK  | Düzenleyici Denetleyici Kurum     | Regulatory Supervisory Authority     |
| FK   | Fon Kurucu - Temsilci             | Fund Founder - Representative        |
| BDK  | Bağımsız Denetim Kuruluşu        | Independent Audit Firm               |
| DCS  | Derecelendirme Şirketi            | Rating Agency                        |
| DS   | Değerlendirme Şirketi             | Valuation Company                    |
| DG   | Diğer                             | Other                                |

#### Example Response

```json
[
  {
    "id": "5900",
    "title": "1000 YATIRIMLAR HOLDİNG A.Ş.",
    "stockCode": "BINHO",
    "memberType": "IGS",
    "kfifUrl": "https://www.kap.org.tr/tr/kfif/8acae2c48b2fa25a018bba0a5034596d"
  },
  {
    "id": "2501",
    "title": "24 GAYRİMENKUL VE GİRİŞİM SERMAYESİ PORTFÖY YÖNETİMİ A.Ş.",
    "stockCode": "YGP",
    "memberType": "FK, PYS"
  }
]
```

---

### Company Securities

Returns summary information for listed companies (IGS type) along with
general information for all securities belonging to each company.

| Property | Value                                                                  |
|----------|------------------------------------------------------------------------|
| URL      | `/api/vyk/memberSecurities`                                            |
| Method   | `GET`                                                                  |
| Version  | 1.0                                                                    |
| Host     | `https://apigwdev.mkk.com.tr`                                         |

#### Response Parameters

Returns a `MemberSecuritiesResponse` containing:

| Field        | Type   | Description                                    |
|--------------|--------|------------------------------------------------|
| `member`     | Object | Company info (CompanyInfo)                     |
| `securities` | List   | Securities list with the following fields      |

**Security fields:**

| Field              | Type   | Description                                  |
|--------------------|--------|----------------------------------------------|
| `isin`             | String | ISIN code                                    |
| `isinDesc`         | String | ISIN description                             |
| `borsaKodu`        | String | Stock exchange code                          |
| `takasKodu`        | String | Clearing code                                |
| `tertipGroup`      | String | Issue group                                  |
| `capital`          | String | Capital                                      |
| `currentCapital`   | String | Current capital                              |
| `groupCode`        | String | Group code                                   |
| `groupCodeDesc`    | String | Group code description                       |
| `borsadaIslemeAcik`| String | Whether trading is active on the exchange    |

---

### Company Detail

Returns general information for a specific company.

| Property | Value                                                                  |
|----------|------------------------------------------------------------------------|
| URL      | `/api/vyk/memberDetail/{id}`                                           |
| Method   | `GET`                                                                  |
| Version  | 1.0                                                                    |
| Host     | `https://apigwdev.mkk.com.tr`                                         |

#### Request Parameters

| Field | Type    | Description                               | Required | Format             |
|-------|---------|-------------------------------------------|----------|--------------------|
| `id`  | Integer | Company ID (path parameter)               | Yes      | `^\\d{1,5}$`       |

#### Response Parameters

Returns a `CompanyDetailResponse`:

| Field              | Type   | Description                              |
|--------------------|--------|------------------------------------------|
| `nameTr`           | String | Turkish field name                       |
| `nameEn`           | String | English field name                       |
| `key`              | String | Field key identifier                     |
| `publishDateTime`  | String | Publication date/time                    |
| `value`            | Object | Field value (structure varies by field)  |

---

### Fund List

Returns the list of all funds.

| Property | Value                                                                  |
|----------|------------------------------------------------------------------------|
| URL      | `/api/vyk/funds`                                                       |
| Method   | `GET`                                                                  |
| Version  | 1.0                                                                    |
| Host     | `https://apigwdev.mkk.com.tr`                                         |

#### Query Parameters (Optional Filters)

| Field       | Type  | Description         |
|-------------|-------|---------------------|
| `fundState` | Array | Filter by fund state |
| `fundClass` | Array | Filter by fund class |
| `fundType`  | Array | Filter by fund type  |

#### Response Parameters

| Field              | Type    | Description                                           | Required |
|--------------------|---------|-------------------------------------------------------|----------|
| `fundId`           | Integer | Unique fund ID                                        | Yes      |
| `fundName`         | String  | Fund title                                            | Yes      |
| `fundCode`         | String  | Fund code                                             | Yes      |
| `fundType`         | String  | Fund type (see table below)                           | Yes      |
| `fundClass`        | String  | Fund class (see table below)                          | Yes      |
| `fundExpiry`       | String  | Fund maturity type: `VL` (Fixed-term), `VS` (Open-ended) | Yes  |
| `fundState`        | String  | Fund status: `Y` (Active), `N` (Passive), `T` (Liquidation) | Yes |
| `title`            | String  | Fund founder title                                    | Yes      |
| `umbMemberTypes`   | String  | Umbrella fund founder type                            | Yes      |
| `fundMemberTypes`  | String  | Fund founder type                                     | Yes      |
| `kapUrl`           | String  | Fund detail page URL on KAP                           | Yes      |
| `nonInactiveCount` | Integer | Number of types linked to the fund's founder          | Yes      |
| `fundCompanyId`    | Integer | Fund founder company ID                               | Yes      |
| `fundCompanyTitle` | String  | Fund founder company title                            | Yes      |

**Fund Types:**

| Code | Turkish                                    | English                             |
|------|--------------------------------------------|-------------------------------------|
| SYF  | Şemsiye Yatırım Fonu (YF)                 | Umbrella Investment Fund            |
| KGF  | Koruma Amaçlı - Garantili Şemsiye YF      | Capital Protected - Guaranteed Fund |
| EYF  | Emeklilik Yatırım Fonu (EYF)              | Pension Investment Fund             |
| OKS  | OKS Emeklilik Yatırım Fonu                | Auto-Enrollment Pension Fund        |
| YYF  | Yabancı Yatırım Fonu (YYF)                | Foreign Investment Fund             |
| BYF  | Borsa Yatırım Fonu (BYF)                  | Exchange Traded Fund (ETF)          |
| VFF  | Varlık Finansman Fonları (VFF)             | Asset Finance Funds                 |
| KFF  | Konut Finansman Fonları (KFF)              | Housing Finance Funds               |
| GMF  | Gayrimenkul Yatırım Fonları (GMF)          | Real Estate Investment Funds        |
| GSF  | Girişim Sermayesi Yatırım Fonu (GSF)       | Venture Capital Investment Fund     |
| PFF  | Proje Finansman Fonu (PFF)                 | Project Finance Fund                |

**Fund Classes:**

| Code | Turkish                  | English              |
|------|--------------------------|----------------------|
| DG   | Diğer                    | Other                |
| PFF  | Proje Finansman Fonu     | Project Finance Fund |
| KTF  | Katılım Fonu             | Participation Fund   |
| HS   | Hisse Yoğun              | Equity Heavy         |
| SF   | Serbest Fon              | Hedge Fund           |

#### Example Response

```json
[
  {
    "fundId": 3929,
    "fundName": "AK PORTFÖY BİRİNCİ KAMU BORÇLANMA ARAÇLARI (TL) ÖZEL FONU",
    "fundCode": "CVE",
    "fundType": "YF",
    "fundClass": "DG",
    "fundExpiry": "VS",
    "fundState": "Y",
    "umbMemberTypes": "FK,PYS",
    "fundMemberTypes": "FK,PYS",
    "kapUrl": "https://www.kap.org.tr/tr/fon-bilgileri/genel/cve-ak-portfoy-birinci-kamu-borclanma-araclari-tl-ozel-fonu",
    "nonInactiveCount": 2,
    "fundCompanyId": "2283",
    "fundCompanyTitle": "AK PORTFÖY YÖNETİMİ A.Ş."
  }
]
```

---

### Fund Detail

Returns general information for a specific fund.

| Property | Value                                                                  |
|----------|------------------------------------------------------------------------|
| URL      | `/api/vyk/fundDetail/{id}`                                             |
| Method   | `GET`                                                                  |
| Version  | 1.0                                                                    |
| Host     | `https://apigwdev.mkk.com.tr`                                         |

#### Request Parameters

| Field    | Type    | Description                          | Required | Format  |
|----------|---------|--------------------------------------|----------|---------|
| `fundId` | Integer | Fund unique ID (path parameter)      | Yes      | `[0-9]` |

**Test data note:** Fund IDs 4282, 4320, and 4372 can be used for test
queries. Since test data is constructed from production data for a specific
historical period, some fund detail fields may have `null` values due to
updates outside that period.

#### Example Request

```text
GET /api/vyk/fundDetail/4282
```

#### Response Parameters

The response varies by fund type. It returns an array of detail sections:

| Field              | Type   | Description                              |
|--------------------|--------|------------------------------------------|
| `nameTr`           | String | Turkish field name                       |
| `nameEn`           | String | English field name                       |
| `key`              | String | Field key identifier                     |
| `publishDateTime`  | String | Publication date/time                    |
| `value`            | Object | Field value (structure varies by field)  |
| `codeKey`          | String | Code key (present in some sections)      |

**Common detail sections include:**

| Key                                       | English Name                                              |
|-------------------------------------------|-----------------------------------------------------------|
| `kpy81_acc2_fon_ucret_kom_info`           | Fund Management Fee Rates and Purchase Redemption Commissions |
| `kpy81_acc3_fon_yonetici`                 | Fund Portfolio Managers                                   |
| `kpy81_acc4_ilet_adres_tel_fax`           | Address, Phone and Fax                                    |
| `kpy81_acc10_diger_hususlar`              | Miscellaneous                                             |
| `kpy81_acc1_fon_sem_unvan`                | Umbrella Fund Title                                       |
| `kpy81_acc1_fon_sem_tur`                  | Category of Umbrella Fund                                 |
| `kpy81_acc1_kurucu_unvan`                 | Title of Founder                                          |
| `kpy81_acc1_portfoy_yon_kurulus`          | Portfolio Manager Institution                             |
| `kpy81_acc1_ISIN`                         | ISIN Code                                                 |
| `kpy81_acc1_halka_arz2`                   | Public Offering Date of Fund                              |
| `kpy81_acc1_fon_sure`                     | Duration of Fund                                          |
| `kpy81_acc1_fon_tasfiye_tarih`            | Liquidation Date of Fund                                  |
| `kpy81_acc1_temel_alim_satim_info`        | Main Trading Information                                  |
| `kpy81_acc1_fon_portfoy_info`             | Information about Fund Portfolio                          |
| `kpy81_acc1_fon_karsilastirma`            | Fund Benchmark                                            |
| `kpy81_acc4_yetkili`                      | Contact People                                            |
| `kpy81_acc1_bdk`                          | Independent Audit Company                                 |
| `kpy81_acc1_fon_kar_dagitim_esaslari`     | Transfer of the Fund's Net Income to Unit Holders         |
| `kpy81_acc1_fon_kamuyu_aydinlatma_esaslari`| Principles of Disclosure                                 |
| `kpy81_acc1_esik_deger1`                  | Threshold Value                                           |
| `kpy81_acc3_fon_mudur`                    | Fund Manager                                              |
| `kpy81_acc1_amac_strateji`                | Investment Strategy and Risk Indicator of Fund            |

#### Example Response (partial)

```json
[
  {
    "nameTr": "ISIN Kodu",
    "nameEn": "ISIN Code",
    "key": "kpy81_acc1_ISIN",
    "publishDateTime": "18/01/2023 18:13:40",
    "value": "TRYISPO01108"
  },
  {
    "nameTr": "Kurucunun Unvanı",
    "nameEn": "Title of Founder",
    "key": "kpy81_acc1_kurucu_unvan",
    "publishDateTime": null,
    "value": "İŞ PORTFÖY YÖNETİMİ A.Ş."
  },
  {
    "nameTr": "Bağımsız Denetim Kuruluşu",
    "nameEn": "Independent Audit Company",
    "key": "kpy81_acc1_bdk",
    "publishDateTime": null,
    "value": "PwC BAĞIMSIZ DENETİM VE SERBEST MUHASEBECİ MALİ MÜŞAVİRLİK A.Ş",
    "codeKey": "92"
  }
]
```

## Data Notes

- Data retrieved through the APIs consists of historical production data. As a
  result, disclosures, company lists, fund lists, and detail information may
  differ from the current state in the live KAP system.
- Disclosure indices start from **538004**. Disclosures between 84196-538004
  are pre-KAP 4.0 and are only available in HTML format.
- The test environment disclosure range is **1091689 to 1231017**.
- Attachment URLs in disclosure detail responses are not publicly accessible.
  Download attachments via the API and store them in your own system before
  serving them to end users.
