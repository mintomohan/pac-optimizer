// This is a single-line comment

function FindProxyForURL(url, host) {
    
    // Check if the request is for localhost
    if (shExpMatch(host, "localhost")) {
        return "DIRECT"; // Localhost access
    }

    if (dnsDomainIs(host, ".example.com")) { // Example domain
        return "PROXY proxy.example.com:8080";
    }

    var test = "http://test.com"; // URL inside string

    var tricky = "this // is not a comment";  // Should not strip inside string

    :this line starts with a colon and should not be removed
    /this line starts with a slash and might look like a comment

    // More comments

    
    return "DIRECT";
}
