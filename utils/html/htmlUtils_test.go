package html

import (
	"strings"
	"testing"
)

func TestGetLinks(t *testing.T) {
	t.Run("Receive links from the html successfully", func(t *testing.T) {

		const htmlContent = `<html> <a href="http://test-url/1"> <a href="http://test-url/2"> </html>`
		body := strings.NewReader(htmlContent)

		linkMap := GetLinks(body)
		if len(linkMap) != 2 {
			t.Errorf("Expected to return size of linkMap as 2 but got %v", len(linkMap))
		}
		if !linkMap["http://test-url/1"] || !linkMap["http://test-url/2"] {
			t.Errorf("Expected linkMap to have urls but got linkMap entry is empty for either of the urls")
		}
	})

	t.Run("Receive links from the sample page successfully", func(t *testing.T) {

		const htmlContent = `<html>
<!-- Text between angle brackets is an HTML tag and is not displayed.
Most tags, such as the HTML and /HTML tags that surround the contents of
a page, come in pairs; some tags, like HR, for a horizontal rule, stand 
alone. Comments, such as the text you're reading, are not displayed when
the Web page is shown. The information between the HEAD and /HEAD tags is 
not displayed. The information between the BODY and /BODY tags is displayed.-->
<head>
<title>Enter a title, displayed at the top of the window.</title>
</head>
<!-- The information between the BODY and /BODY tags is displayed.-->
<body>
<h1>Enter the main heading, usually the same as the title.</h1>
<p>Be <b>bold</b> in stating your key points. Put them in a list: </p>
<ul>
<li>The first item in your list</li>
<li>The second item; <i>italicize</i> key words</li>
</ul>
<p>Improve your image by including an image. </p>
<p><img src="http://www.mygifs.com/CoverImage.gif" alt="A Great HTML Resource"></p>
<p>Add a link to your favorite <a href="https://www.dummies.com/">Web site</a>.
Break up your page with a horizontal rule or two. </p>
<hr>
<p>Finally, link to <a href="page2.html">another page</a> in your own Web site.</p>
<!-- And add a copyright notice.-->
<p>&#169; Wiley Publishing, 2011</p>
</body>
</html>`
		body := strings.NewReader(htmlContent)

		linkMap := GetLinks(body)
		if len(linkMap) != 2 {
			t.Errorf("Expected to return size of linkMap as 2 but got %v", len(linkMap))
		}
		if !linkMap["https://www.dummies.com/"] || !linkMap["page2.html"] {
			t.Errorf("Expected linkMap to have urls but got linkMap entry is empty for either of the urls")
		}
	})
}
