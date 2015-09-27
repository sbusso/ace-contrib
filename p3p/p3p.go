package p3p

import (
	"github.com/plimble/ace"
)

// var (
// 	defaultPolicyRef = "http://cdn.turfmedia.com/p3p.xml"
// 	defaultPolicy = "NON DSP COR NID OUR DEL SAM OTR UNR COM NAV INT DEM CNT STA PRE LOC OTC"
// )

// Options stores configurations
type Options struct {
	Policy     string
	PolicyRef  string
}

// Middleware sets CORS headers for every request
// func P3p(options Options) ace.HandlerFunc {
func P3p() ace.HandlerFunc {
	// if options.Policy == nil {
	// 	options.Policy = defaultPolicy
	// }
  //
	// if options.PolicyRef == nil {
	// 	options.PolicyRef = defaultPolicyRef
	// }

	return func(c *ace.C) {
		// origin := c.Request.Header.Get("Origin")
		// requestMethod := c.Request.Header.Get("Access-Control-Request-Method")
		// requestHeaders := c.Request.Header.Get("Access-Control-Request-Headers")
		// P3P: CP="'.owa_coreAPI::getSetting('base', 'p3p_policy').'"'
		c.Writer.Header().Set("P3P", "policyref='http://cdn.turfmedia.com/w3c/p3p.xml', CP='NON DSP COR NID OUR DEL SAM OTR UNR COM NAV INT DEM CNT STA PRE LOC OTC'")

		c.Next()
	}
}


// P3P: policyref="/w3c/p3p.xml", CP="IDC DSP COR CURa ADM DEV TAI CON HIS OUR IND ONL UNI PUR COM NAV INT DEM CNT STA PRE LOC"
