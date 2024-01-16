package main

type DataSource struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
	URL string `json:"URL"`
}

type VendorSeverity struct {
	NVD        int `json:"nvd"`
	Redhat     int `json:"redhat"`
	Ubuntu     int `json:"ubuntu"`
}

type CVSS struct {
	NVD struct {
		V3Vector string `json:"V3Vector"`
		V3Score float64 `json:"V3Score"`
	} `json:"nvd"`
	Redhat struct {
		V3Vector string `json:"V3Vector"`
		V3Score float64 `json:"V3Score"`
	} `json:"redhat"`
}

type Vulnerability struct {
	VulnerabilityID    string        `json:"VulnerabilityID"`
	PkgID              string        `json:"PkgID"`
	PkgName            string        `json:"PkgName"`
	InstalledVersion   string        `json:"InstalledVersion"`
	FixedVersion       string        `json:"FixedVersion"`
	Status             string        `json:"Status"`
	Layer              map[string]string `json:"Layer"`
	SeveritySource     string        `json:"SeveritySource"`
	PrimaryURL         string        `json:"PrimaryURL"`
	DataSource         DataSource    `json:"DataSource"`
	Title              string        `json:"Title"`
	Description        string        `json:"Description"`
	Severity           string        `json:"Severity"`
	VendorSeverity     VendorSeverity `json:"VendorSeverity"`
	CVSS               CVSS          `json:"CVSS"`
	References         []string      `json:"References"`
	PublishedDate      string        `json:"PublishedDate"`
	LastModifiedDate   string        `json:"LastModifiedDate"`
}

type Panzerformat struct{
        VulnerabilityID    string        `json:"VulnerabilityID"`
        SeveritySource     string        `json:"SeveritySource"`
        LastModifiedDate   string        `json:"LastModifiedDate"`
        PublishedDate      string        `json:"PublishedDate"`
        Severity           string        `json:"Severity"`
        PkgName            string        `json:"PkgName"`
        InstalledVersion   string        `json:"InstalledVersion"`
        FixedVersion       string        `json:"FixedVersion"`
}
