package main

import (
	"fmt"
	"testing"
)

func Test_getParsedLog(t *testing.T) {
	remarks := []string{
		`time="2025-11-11T15:50:08+05:30" level=trace msg="f is not nil &os.File{file:(*os.file)(0xc00013e000)}"`,
		`time="2025-11-11T15:50:08+05:30" level=trace msg="connecting db"`,
		`time="2025-11-11T15:50:08+05:30" level=trace msg="db connected."`,
		`time="2025-11-11T15:50:08+05:30" level=trace msg="processing app digitags vincotte"`,
		`time="2025-11-11T15:50:08+05:30" level=trace msg="exproting 4 assets "`,
		`time="2025-11-11T15:50:10+05:30" level=trace msg="_____ crepo store.CachedClient{ID:\"TEST123\", Name:\"TEST\", Language:\"\", Address1:\"\", Street:\"\", City:\"\", PostalCode:\"\"}"`,
		`time="2025-11-11T15:50:11+05:30" level=trace msg="_______ returjing "`,
		`time="2025-11-11T15:50:12+05:30" level=trace msg="exproting 4 assets "`,
		`time="2025-11-11T15:50:12+05:30" level=trace msg="exproting 36 CharacteristicsValuesRepo "`,
		`time="2025-11-11T15:50:12+05:30" level=trace msg="exproting 4 DigitagsWorkOrder "`,
		`time="2025-11-11T15:50:12+05:30" level=trace msg=" cache 4 assetsfor client TEST123"`,
		`time="2025-11-11T15:50:12+05:30" level=trace msg="exproting 0 Characteristics "`,
		`time="2025-11-11T15:50:14+05:30" level=trace msg="inserting asset ICIM-PDC-000089 : TEST123"`,
		`time="2025-11-11T15:50:15+05:30" level=trace msg="inserting asset 011-1W0496 : TEST123"`,
		`time="2025-11-11T15:50:16+05:30" level=trace msg="inserting asset 041-K115-07 : TEST123"`,
		`time="2025-11-11T15:50:16+05:30" level=trace msg="inserting asset 037-161-25 : TEST123"`,
		`time="2025-11-11T15:50:17+05:30" level=debug msg="upserting control history" count=4 index=3`,
		`time="2025-11-11T15:50:17+05:30" level=trace msg=" control history query \n\t\tINSERT INTO ControlHistory \n\t\t\t(id, label, visitdate, validity, controltype, assetid, clientid, decisionen, decisionfi, decisionnl, decisionfr,inspector, reportid) \n\t\tVALUES \n\t\t\t(\"443944ee5e36022f9077bfd89cb72103\", \"1\", \"2020-12-21 00:00:00 +0000 UTC\", \"2099-12-31 00:00:00 +0000 UTC\",\"\", \"ICIM-PDC-000089\",\"TEST123\", \"\", \"\", \"\", \"\", \"\", \"\"), (\"33d786cf15ab1e9f06f15dc9bf133da8\", \"1\", \"2021-11-10 00:00:00 +0000 UTC\", \"2099-12-31 00:00:00 +0000 UTC\",\"\", \"011-1W0496\",\"TEST123\", \"\", \"\", \"\", \"\", \"\", \"\"), (\"9d086c53352458c6c62a46af3c0cff3a\", \"1\", \"2025-08-27 00:00:00 +0000 UTC\", \"2099-12-31 00:00:00 +0000 UTC\",\"\", \"041-K115-07\",\"TEST123\", \"\", \"\", \"\", \"\", \"\", \"\"), (\"1fc5f660f12dec88ce1432cd26c45257\", \"1\", \"2025-02-25 00:00:00 +0000 UTC\", \"2099-12-31 00:00:00 +0000 UTC\",\"\", \"037-161-25\",\"TEST123\", \"\", \"\", \"\", \"\", \"\", \"\")\n\t\tas new\n\t\tON DUPLICATE KEY UPDATE \n\t\t\tlabel = new.label,\n\t\t\tvisitdate= new.visitdate,\n\t\t\tvalidity= new.validity,\n\t\t\tdecisionen= new.decisionen,\n\t\t\tdecisionnl= new.decisionnl,\n\t\t\tdecisionfr= new.decisionfr,\n\t\t\tdecisionfi= new.decisionfi,\n\t\t\tcontroltype= new.controltype,\n\t\t\tinspector= new.inspector,\n\t\t\treportid= new.reportid"`,
		`time="2025-11-11T15:50:18+05:30" level=debug msg="4 rows affected while upserting control history"`,
		`time="2025-11-11T15:50:18+05:30" level=trace msg="exproting 0 ControlHistoryRemarksRepo "`,
		`time="2025-09-30T15:02:40+05:30" level=error msg="error in processing csv &errors.errorString{s:\"error getting file asdfasdf.pdf: &fs.PathError{Op:\\\"open\\\", Path:\\\"reg/asdfasd.pdf\\\", Err:0x2}\"}"`,
	}
	inChan := make(chan string, 1000)
	p := newProcessor(inChan, nil)
	for _, item := range remarks {
		fmt.Printf("%#v", p.getParsedLog(item))
	}
	fmt.Println()
}
