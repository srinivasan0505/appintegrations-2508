/**
 * MCP Server function for <p>Creates and persists a DataIntegration resource.</p> <note> <p>You cannot create a DataIntegration association for a DataIntegration that has been previously associated. Use a different DataIntegration, or recreate the DataIntegration using the <code>CreateDataIntegration</code> API.</p> </note>
 */

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.*;
import java.util.function.Function;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;

class Post_Data_IntegrationsMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_Data_IntegrationsHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("X-Amz-Content-Sha256")) {
            queryParams.add("X-Amz-Content-Sha256=" + args.get("X-Amz-Content-Sha256"));
        }
        if (args.containsKey("X-Amz-Date")) {
            queryParams.add("X-Amz-Date=" + args.get("X-Amz-Date"));
        }
        if (args.containsKey("X-Amz-Algorithm")) {
            queryParams.add("X-Amz-Algorithm=" + args.get("X-Amz-Algorithm"));
        }
        if (args.containsKey("X-Amz-Credential")) {
            queryParams.add("X-Amz-Credential=" + args.get("X-Amz-Credential"));
        }
        if (args.containsKey("X-Amz-Security-Token")) {
            queryParams.add("X-Amz-Security-Token=" + args.get("X-Amz-Security-Token"));
        }
        if (args.containsKey("X-Amz-Signature")) {
            queryParams.add("X-Amz-Signature=" + args.get("X-Amz-Signature"));
        }
        if (args.containsKey("X-Amz-SignedHeaders")) {
            queryParams.add("X-Amz-SignedHeaders=" + args.get("X-Amz-SignedHeaders"));
        }
        if (args.containsKey("Name")) {
            queryParams.add("Name=" + args.get("Name"));
        }
        if (args.containsKey("ClientToken")) {
            queryParams.add("ClientToken=" + args.get("ClientToken"));
        }
        if (args.containsKey("Description")) {
            queryParams.add("Description=" + args.get("Description"));
        }
        if (args.containsKey("SourceURI")) {
            queryParams.add("SourceURI=" + args.get("SourceURI"));
        }
        if (args.containsKey("KmsKey")) {
            queryParams.add("KmsKey=" + args.get("KmsKey"));
        }
        if (args.containsKey("ObjectConfiguration")) {
            queryParams.add("ObjectConfiguration=" + args.get("ObjectConfiguration"));
        }
        if (args.containsKey("ScheduleConfig")) {
            queryParams.add("ScheduleConfig=" + args.get("ScheduleConfig"));
        }
        if (args.containsKey("Tags")) {
            queryParams.add("Tags=" + args.get("Tags"));
        }
        if (args.containsKey("FileConfiguration")) {
            queryParams.add("FileConfiguration=" + args.get("FileConfiguration"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_data_integrations" + queryString;
                
                HttpClient client = HttpClient.newHttpClient();
                HttpRequest httpRequest = HttpRequest.newBuilder()
                    .uri(URI.create(url))
                    .header("Authorization", "Bearer " + config.getApiKey())
                    .header("Accept", "application/json")
                    .GET()
                    .build();
                
                HttpResponse<String> response = client.send(httpRequest, HttpResponse.BodyHandlers.ofString());
                
                if (response.statusCode() >= 400) {
                    return new MCPServer.MCPToolResult("API error: " + response.body(), true);
                }
                
                // Pretty print JSON
                ObjectMapper mapper = new ObjectMapper();
                JsonNode jsonNode = mapper.readTree(response.body());
                String prettyJson = mapper.writerWithDefaultPrettyPrinter().writeValueAsString(jsonNode);
                
                return new MCPServer.MCPToolResult(prettyJson);
                
            } catch (IOException | InterruptedException e) {
                return new MCPServer.MCPToolResult("Request failed: " + e.getMessage(), true);
            } catch (Exception e) {
                return new MCPServer.MCPToolResult("Unexpected error: " + e.getMessage(), true);
            }
        };
    }
    
    public static MCPServer.Tool createPost_Data_IntegrationsTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> X-Amz-Content-Sha256Property = new HashMap<>();
        X-Amz-Content-Sha256Property.put("type", "string");
        X-Amz-Content-Sha256Property.put("required", false);
        X-Amz-Content-Sha256Property.put("description", "");
        properties.put("X-Amz-Content-Sha256", X-Amz-Content-Sha256Property);
        Map<String, Object> X-Amz-DateProperty = new HashMap<>();
        X-Amz-DateProperty.put("type", "string");
        X-Amz-DateProperty.put("required", false);
        X-Amz-DateProperty.put("description", "");
        properties.put("X-Amz-Date", X-Amz-DateProperty);
        Map<String, Object> X-Amz-AlgorithmProperty = new HashMap<>();
        X-Amz-AlgorithmProperty.put("type", "string");
        X-Amz-AlgorithmProperty.put("required", false);
        X-Amz-AlgorithmProperty.put("description", "");
        properties.put("X-Amz-Algorithm", X-Amz-AlgorithmProperty);
        Map<String, Object> X-Amz-CredentialProperty = new HashMap<>();
        X-Amz-CredentialProperty.put("type", "string");
        X-Amz-CredentialProperty.put("required", false);
        X-Amz-CredentialProperty.put("description", "");
        properties.put("X-Amz-Credential", X-Amz-CredentialProperty);
        Map<String, Object> X-Amz-Security-TokenProperty = new HashMap<>();
        X-Amz-Security-TokenProperty.put("type", "string");
        X-Amz-Security-TokenProperty.put("required", false);
        X-Amz-Security-TokenProperty.put("description", "");
        properties.put("X-Amz-Security-Token", X-Amz-Security-TokenProperty);
        Map<String, Object> X-Amz-SignatureProperty = new HashMap<>();
        X-Amz-SignatureProperty.put("type", "string");
        X-Amz-SignatureProperty.put("required", false);
        X-Amz-SignatureProperty.put("description", "");
        properties.put("X-Amz-Signature", X-Amz-SignatureProperty);
        Map<String, Object> X-Amz-SignedHeadersProperty = new HashMap<>();
        X-Amz-SignedHeadersProperty.put("type", "string");
        X-Amz-SignedHeadersProperty.put("required", false);
        X-Amz-SignedHeadersProperty.put("description", "");
        properties.put("X-Amz-SignedHeaders", X-Amz-SignedHeadersProperty);
        Map<String, Object> NameProperty = new HashMap<>();
        NameProperty.put("type", "string");
        NameProperty.put("required", true);
        NameProperty.put("description", "Input parameter: The name of the DataIntegration.");
        properties.put("Name", NameProperty);
        Map<String, Object> ClientTokenProperty = new HashMap<>();
        ClientTokenProperty.put("type", "string");
        ClientTokenProperty.put("required", false);
        ClientTokenProperty.put("description", "Input parameter: A unique, case-sensitive identifier that you provide to ensure the idempotency of the request. If not provided, the Amazon Web Services SDK populates this field. For more information about idempotency, see <a href=\"https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/\">Making retries safe with idempotent APIs</a>.");
        properties.put("ClientToken", ClientTokenProperty);
        Map<String, Object> DescriptionProperty = new HashMap<>();
        DescriptionProperty.put("type", "string");
        DescriptionProperty.put("required", false);
        DescriptionProperty.put("description", "Input parameter: A description of the DataIntegration.");
        properties.put("Description", DescriptionProperty);
        Map<String, Object> SourceURIProperty = new HashMap<>();
        SourceURIProperty.put("type", "string");
        SourceURIProperty.put("required", true);
        SourceURIProperty.put("description", "Input parameter: The URI of the data source.");
        properties.put("SourceURI", SourceURIProperty);
        Map<String, Object> KmsKeyProperty = new HashMap<>();
        KmsKeyProperty.put("type", "string");
        KmsKeyProperty.put("required", true);
        KmsKeyProperty.put("description", "Input parameter: The KMS key for the DataIntegration.");
        properties.put("KmsKey", KmsKeyProperty);
        Map<String, Object> ObjectConfigurationProperty = new HashMap<>();
        ObjectConfigurationProperty.put("type", "string");
        ObjectConfigurationProperty.put("required", false);
        ObjectConfigurationProperty.put("description", "Input parameter: The configuration for what data should be pulled from the source.");
        properties.put("ObjectConfiguration", ObjectConfigurationProperty);
        Map<String, Object> ScheduleConfigProperty = new HashMap<>();
        ScheduleConfigProperty.put("type", "string");
        ScheduleConfigProperty.put("required", true);
        ScheduleConfigProperty.put("description", "Input parameter: The name of the data and how often it should be pulled from the source.");
        properties.put("ScheduleConfig", ScheduleConfigProperty);
        Map<String, Object> TagsProperty = new HashMap<>();
        TagsProperty.put("type", "string");
        TagsProperty.put("required", false);
        TagsProperty.put("description", "Input parameter: The tags used to organize, track, or control access for this resource. For example, { \"tags\": {\"key1\":\"value1\", \"key2\":\"value2\"} }.");
        properties.put("Tags", TagsProperty);
        Map<String, Object> FileConfigurationProperty = new HashMap<>();
        FileConfigurationProperty.put("type", "string");
        FileConfigurationProperty.put("required", false);
        FileConfigurationProperty.put("description", "Input parameter: The configuration for what files should be pulled from the source.");
        properties.put("FileConfiguration", FileConfigurationProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_data_integrations",
            "<p>Creates and persists a DataIntegration resource.</p> <note> <p>You cannot create a DataIntegration association for a DataIntegration that has been previously associated. Use a different DataIntegration, or recreate the DataIntegration using the <code>CreateDataIntegration</code> API.</p> </note>",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_Data_IntegrationsHandler(config));
    }
    
}