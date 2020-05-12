package main

import (
	"fmt"
	"github.com/mmcdole/gofeed/atom"
	"strings"
)

var feedData = `
     <?xml version="1.0" encoding="UTF-8"?>
     <feed xmlns="https://www.w3.org/2005/Atom"
         xmlns:rolie="urn:ietf:params:xml:ns:rolie-1.0"
         xml:lang="en-US">
       <id>2a7e265a-39bc-43f2-b711-b8fd9264b5c9</id>
       <title type="text">
           Atom-formatted representation of
           a Feed of XML vulnerability documents
       </title>
       <category
           scheme="urn:ietf:params:rolie:category:information-type"
           term="vulnerability"/>
       <updated>2016-05-04T18:13:51.0Z</updated>
       <link rel="self"
           href="https://example.org/provider/public/vulns"/>
       <link rel="service"
           href="https://example.org/rolie/servicedocument"/>
       <entry>
         <rolie:format ns="urn:ietf:params:xml:ns:exampleformat"/>
         <id>dd786dba-88e6-440b-9158-b8fae67ef67c</id>
         <title>Sample Vulnerability</title>
         <published>2015-08-04T18:13:51.0Z</published>
         <updated>2015-08-05T18:13:51.0Z</updated>
         <summary>A vulnerability issue identified by CVE-...</summary>
         <content type="application/xml"
             src="https://example.org/provider/vulns/123456/data"/>
       </entry>
       <entry>
           <!-- ...another entry... -->
       </entry>
     </feed>
`

func main() {

	fp := atom.Parser{}
	atomFeed, _ := fp.Parse(strings.NewReader(feedData))
	fmt.Println(atomFeed.Title)

	rolieEntry := atomFeed.Entries[0]
	rolieFormats := rolieEntry.Extensions["rolie"]["format"]
	rolieFormatCopy := rolieFormats[0]
	rolieFormats = append(rolieFormats, rolieFormatCopy)
}
