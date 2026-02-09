package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	kap "github.com/knckknckknck/kap-go"
)

func main() {
	apiKey := os.Getenv("KAP_API_KEY")
	secretKey := os.Getenv("KAP_SECRET_KEY")

	if apiKey == "" || secretKey == "" {
		log.Fatal("KAP_API_KEY and KAP_SECRET_KEY environment variables are required")
	}

	client := kap.NewClient(apiKey,
		kap.WithBasicAuth(apiKey, secretKey),
		kap.WithTimeout(30*time.Second),
	)

	ctx := context.Background()

	// 1. Last Disclosure Index
	fmt.Println("=== 1. LastDisclosureIndex ===")
	lastIndex, err := client.LastDisclosureIndex(ctx)
	if err != nil {
		log.Fatalf("  FAIL: %v", err)
	}
	fmt.Printf("  Last disclosure index: %s\n\n", lastIndex)

	// 2. Disclosures (start from test range)
	fmt.Println("=== 2. Disclosures ===")
	disclosures, err := client.Disclosures(ctx, 1092228, nil)
	if err != nil {
		log.Fatalf("  FAIL: %v", err)
	}
	fmt.Printf("  Fetched %d disclosures\n", len(disclosures))
	for i, d := range disclosures {
		fmt.Printf("    [%s] %s (type: %s, class: %s)\n",
			d.DisclosureIndex, d.Title, d.DisclosureType, d.DisclosureClass)
		if i >= 4 {
			fmt.Printf("    ... and %d more\n", len(disclosures)-5)
			break
		}
	}
	fmt.Println()

	// 3. Disclosures with filters
	fmt.Println("=== 3. Disclosures (filtered) ===")
	filtered, err := client.Disclosures(ctx, 1092228, &kap.DisclosureListParams{
		DisclosureClass: "DG",
		DisclosureType:  "DG",
	})
	if err != nil {
		log.Fatalf("  FAIL: %v", err)
	}
	fmt.Printf("  Fetched %d filtered disclosures (class=DG, type=DG)\n", len(filtered))
	for i, d := range filtered {
		fmt.Printf("    [%s] %s\n", d.DisclosureIndex, d.Title)
		if i >= 2 {
			fmt.Printf("    ... and %d more\n", len(filtered)-3)
			break
		}
	}
	fmt.Println()

	// 4. Disclosure Detail
	fmt.Println("=== 4. DisclosureDetail ===")
	detail, err := client.DisclosureDetail(ctx, 1211180, "data", "")
	if err != nil {
		log.Fatalf("  FAIL: %v", err)
	}
	fmt.Printf("  Index: %s\n", detail.DisclosureIndex)
	fmt.Printf("  Sender: %s\n", detail.SenderTitle)
	fmt.Printf("  Type: %s, Class: %s\n", detail.DisclosureType, detail.DisclosureClass)
	fmt.Printf("  Reason: %s\n", detail.DisclosureReason)
	fmt.Printf("  Time: %s\n", detail.Time)
	if detail.Subject.TR != nil {
		fmt.Printf("  Subject (TR): %s\n", *detail.Subject.TR)
	}
	if detail.Subject.EN != nil {
		fmt.Printf("  Subject (EN): %s\n", *detail.Subject.EN)
	}
	fmt.Printf("  Attachments: %d\n", len(detail.AttachmentURLs))
	for _, a := range detail.AttachmentURLs {
		fmt.Printf("    %s\n", a.FileName)
	}
	fmt.Printf("  Presentations: %d\n", len(detail.Presentation))
	fmt.Printf("  FlatData items: %d\n", len(detail.FlatData))
	fmt.Printf("  HTML messages: %d\n", len(detail.HTMLMessages))
	fmt.Println()

	// 5. Download Attachment
	fmt.Println("=== 5. DownloadAttachment ===")
	if len(detail.AttachmentURLs) > 0 {
		attURL := detail.AttachmentURLs[0].URL
		var attID string
		for i := len(attURL) - 1; i >= 0; i-- {
			if attURL[i] == '/' {
				attID = attURL[i+1:]
				break
			}
		}
		body, disposition, err := client.DownloadAttachment(ctx, attID)
		if err != nil {
			fmt.Printf("  ERROR: %v (non-fatal)\n\n", err)
		} else {
			data, _ := io.ReadAll(body)
			body.Close()
			fmt.Printf("  Content-Disposition: %s\n", disposition)
			fmt.Printf("  Downloaded %d bytes\n\n", len(data))
		}
	} else {
		fmt.Println("  Skipped (no attachments)")
		fmt.Println()
	}

	// 6. Blocked Disclosures
	fmt.Println("=== 6. BlockedDisclosures ===")
	blocked, err := client.BlockedDisclosures(ctx)
	if err != nil {
		log.Fatalf("  FAIL: %v", err)
	}
	fmt.Printf("  Response (%d bytes): %.200s", len(blocked), string(blocked))
	if len(blocked) > 200 {
		fmt.Print("...")
	}
	fmt.Println()
	fmt.Println()

	// 7. Members
	fmt.Println("=== 7. Members ===")
	members, err := client.Members(ctx)
	if err != nil {
		log.Fatalf("  FAIL: %v", err)
	}
	fmt.Printf("  Found %d members\n", len(members))
	for i, m := range members {
		fmt.Printf("    %s (ID: %s, stock: %s, type: %s)\n", m.Title, m.ID, m.StockCode, m.MemberType)
		if i >= 2 {
			fmt.Printf("    ... and %d more\n", len(members)-3)
			break
		}
	}
	fmt.Println()

	// 8. Member Securities
	fmt.Println("=== 8. MemberSecurities ===")
	memberSec, err := client.MemberSecurities(ctx)
	if err != nil {
		log.Fatalf("  FAIL: %v", err)
	}
	fmt.Printf("  Found %d entries\n", len(memberSec))
	if len(memberSec) > 0 {
		ms := memberSec[0]
		fmt.Printf("    First: %s (%d securities)\n", ms.Member.SirketUnvan, len(ms.Securities))
		for i, s := range ms.Securities {
			fmt.Printf("      ISIN: %s, BorsaKodu: %s, TakasKodu: %s\n", s.ISIN, s.BorsaKodu, s.TakasKodu)
			if i >= 1 {
				fmt.Printf("      ... and %d more securities\n", len(ms.Securities)-2)
				break
			}
		}
	}
	fmt.Println()

	// 9. Member Detail
	fmt.Println("=== 9. MemberDetail ===")
	if len(members) > 0 {
		var memberID int
		fmt.Sscanf(members[0].ID, "%d", &memberID)
		fields, err := client.MemberDetail(ctx, memberID)
		if err != nil {
			log.Fatalf("  FAIL: %v", err)
		}
		fmt.Printf("  Detail for %s (%d fields):\n", members[0].Title, len(fields))
		for i, f := range fields {
			valStr := truncateRaw(f.Value, 60)
			fmt.Printf("    %s (%s): %s\n", f.NameEN, f.Key, valStr)
			if i >= 4 {
				fmt.Printf("    ... and %d more fields\n", len(fields)-5)
				break
			}
		}
	}
	fmt.Println()

	// 10. Funds
	fmt.Println("=== 10. Funds ===")
	funds, err := client.Funds(ctx, nil)
	if err != nil {
		log.Fatalf("  FAIL: %v", err)
	}
	fmt.Printf("  Found %d funds\n", len(funds))
	for i, f := range funds {
		fmt.Printf("    %s (ID: %d, code: %s, type: %s, state: %s)\n",
			f.FundName, f.FundID, f.FundCode, f.FundType, f.FundState)
		if i >= 2 {
			fmt.Printf("    ... and %d more\n", len(funds)-3)
			break
		}
	}
	fmt.Println()

	// 11. Fund Detail (test fund ID 4282)
	fmt.Println("=== 11. FundDetail ===")
	fundFields, err := client.FundDetail(ctx, 4282)
	if err != nil {
		log.Fatalf("  FAIL: %v", err)
	}
	fmt.Printf("  Detail for fund 4282 (%d fields):\n", len(fundFields))
	for i, f := range fundFields {
		valStr := truncateRaw(f.Value, 60)
		fmt.Printf("    %s (%s): %s\n", f.NameEN, f.Key, valStr)
		if i >= 4 {
			fmt.Printf("    ... and %d more fields\n", len(fundFields)-5)
			break
		}
	}
	fmt.Println()

	// 12. CA Event Status
	fmt.Println("=== 12. CAEventStatus ===")
	caStatus, err := client.CAEventStatus(ctx, "TEST-REF-001")
	if err != nil {
		fmt.Printf("  ERROR: %v (expected — no valid ref ID for test)\n", err)
	} else {
		fmt.Printf("  RefID: %s, Status: %s, Reason: %s\n", caStatus.RefID, caStatus.Status, caStatus.StatusReason)
	}

	fmt.Println("\n=== All 12 endpoint calls completed ===")
}

func truncateRaw(raw json.RawMessage, maxLen int) string {
	if len(raw) == 0 {
		return "<nil>"
	}
	s := string(raw)
	if len(s) > maxLen {
		return s[:maxLen] + "..."
	}
	return s
}
