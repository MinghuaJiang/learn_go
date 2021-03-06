package product_v2

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

type ProductReportFilter struct {
	NameFilter         string `json:"productNAme"`
	ManufacturerFilter string `json:"manufacturer`
	SKUFilter          string `json:"sku"`
}

func reportHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var productFilter ProductReportFilter
		err := json.NewDecoder(r.Body).Decode(&productFilter)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		products, err := searchForProductData(productFilter)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		t := template.New("report_v3.gotmpl").Funcs(template.FuncMap{"mod": func(i, x int) bool { return i%x == 0 }})
		t, err = t.ParseFiles(path.Join("templates", "report_v3.gotmpl"))

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var tmpl bytes.Buffer
		if len(products) > 0 {
			err = t.Execute(&tmpl, products)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		rdr := bytes.NewReader(tmpl.Bytes())

		w.Header().Set("Content-Disposition", "Attachment")
		http.ServeContent(w, r, "report.html", time.Now(), rdr)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
