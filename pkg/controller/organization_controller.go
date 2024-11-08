func setupGrafanaAdmin(ctx context.Context, grafanaURL string, adminUsername string, adminPassword string) (string, error) {
    // Create the admin user if not already created
    // Example: Use HTTP client to send requests to Grafana to create the user
    
    // Generate the API key
    client := &http.Client{}
    data := map[string]string{
        "name": "admin-api-key",
        "role": "Admin",
    }
    jsonData, err := json.Marshal(data)
    if err != nil {
        return "", err
    }

    req, err := http.NewRequest("POST", grafanaURL+"/api/auth/keys", bytes.NewBuffer(jsonData))
    req.SetBasicAuth(adminUsername, adminPassword)
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := ioutil.ReadAll(resp.Body)
        return "", fmt.Errorf("failed to generate API key: %s", string(body))
    }

    var response map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        return "", err
    }

    apiKey, ok := response["key"].(string)
    if !ok {
        return "", fmt.Errorf("failed to retrieve API key from response")
    }

    return apiKey, nil
}
func storeBearerToken(ctx context.Context, c client.Client, namespace string, token string) error {
    secret := &corev1.Secret{
        ObjectMeta: metav1.ObjectMeta{
            Name:      "grafana-secret",
            Namespace: namespace,
        },
        Data: map[string][]byte{
            "bearerToken": []byte(token),
        },
    }
    return c.Create(ctx, secret)
}
func storeBearerToken(ctx context.Context, c client.Client, namespace string, token string) error {
    secret := &corev1.Secret{
        ObjectMeta: metav1.ObjectMeta{
            Name:      "grafana-secret",
            Namespace: namespace,
        },
        Data: map[string][]byte{
            "bearerToken": []byte(token),
        },
    }
    return c.Create(ctx, secret)
}

