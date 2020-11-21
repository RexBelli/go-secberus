# go-secberus

## Example
```
	c, err := NewClientFromEnv()
	if err != nil {
		log.Fatalf("could not create client: %s\n", err)
	}

	orgs, err := c.GetOrganizations()
	if err != nil {
		log.Fatalf("could not get organizations: %s\n", err)
	}

	for _, org := range *orgs {
		fmt.Printf("%+v\n", org)
	}
```

## Debug
This includes code to proxy traffic through an HTTP proxy (such as burp).  Set the DEBUG environment variable to anything to enable it.  `unset` it to turn disable debugging.