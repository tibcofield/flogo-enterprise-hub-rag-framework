Flogo ChatGPT Client Exension
=============================

Please do not use this for your own projects at this stage as it is still in very early stage of development.  If you are a TIBCO customer and would like to use the extension or have feeback then please reach out to the repository owner or your account team. 

Overview
--------

The Flogo ChatGPT extesion acts as a client for the APIs exposed from the ChatGPT developer platform.  It is in the very early days of development and there is a focus on supporting text generation, image generation, embedings creation and storage using the Responses API. 


OpenAI Go API Library Version - v3.29.0

Roadmap
-------

| Client Activity                     | Status                        | Open AI Docs
| ------------------------------------| ------------------------------| ---------------------------------------------------|
| Responses API                       | In Dev for text and Images    | https://platform.openai.com/docs/guides/migrate-to-responses | 
| Images API                          | In Dev                        | https://platform.openai.com/docs/api-reference/images |
| Embedings API                       | Testing                       | https://platform.openai.com/docs/guides/embeddings |
| Files API                           | In Dev                        | https://platform.openai.com/docs/api-reference/files |
| Chat API                            | Out of scope                  ||
| Completions API                     | Out of scope                  | Deprecated and unclear on OpenAI, other platforms direction |
| Realtime API                        | Out of scope                  || 
| Assistants API                      | Out of scope                  ||
| Batch API                           | Out of scope                  ||
| Containers API                      | Out of scope                  ||
| Fine Tunning API                    | Out of scope                  ||
| Graders API                         | Out of scope                  ||
| Moderrations API                    | Out of scope                  ||
| Vector Store API                    | In Scope.                     | Allows you to manage vector stores. |

## ChatGPT Use Cases 

These are some of the example use cases we are being reviewed as part of the client extension development.

| Client Activity                     | Status                        |
| ----------------------------------- | ------------------------------|
| Embeddings                          | Testing                       |
| Responses Text Generation           | In Dev                        |
| Create New Image                    | In Dev                        |
| Edit Existing Image                 | In Dev                        |
| Edit Existing Image Multiple        | In Dev                        |
| Edit Existing Image Mask            | In Dev                        |
| Responses - Structured Response     | Backlog                       |
| Responses - Web Search              | Backlog                       |
| Responses - File Search (RAG)       | Backlog                       |
| App with Flogo MCP                  | Bcaklog                       |     
| Responses - Function Calling        | Out of scope                  |
| Create New Image (mulitple Images)  | Backlog                       |
| Responses - Remote MCP              | Out of scope                  |
| Audio Generation                    | Out of scope                  |
| Deep Research                       | Out of scope                  |  
| Moderation                          | Out of scope                  |

## OpenAI api Go Library

To interact with the open AI API the official go library is used.  More information on the API can be found here:

### SDK Overview
https://github.com/openai/openai-go

### API Docs
https://github.com/openai/openai-go/blob/main/api.md

### Open AI API Overview
https://platform.openai.com/docs/api-reference/introduction


## More Information

For more information please reach out to the repository owner.

## Compatible Embeding Models

|Platform | Model                           | Last Known Test Status |
|---------|---------------------------------|------------------------|
|Open AI  | text-embedding-3-small          | Pass                   | 
|Open AI  | text-embedding-3-large          | Pass                   | 
|Ollama   | nomic-embed-text:latest         | Pass                   | 
|Ollama   | mxbai-embed-large:latest        | Pass                   | 
|Ollama   | bge-m3:latest                   | Pass                   |
|Ollama   | all-minilm:latest               | Pass                   | 
|Ollama   | snowflake-arctic-embed2:latest  | Pass                   |
|Ollama   | bge-large:latest                | Pass                   |
|Ollama   | paraphrase-multilingual:latest  | Pass                   |
|Ollama   | granite-embedding:latest        | Pass                   |
|Ollama   | embeddinggemma:latest           | Pass                   | 

## Default URLs

| Platfrom             | Default Enpoint URL       | Company Website URL   |
|----------------------|---------------------------|---------------------- |
| Open AI.             | https://api.openai.com/v1 | https://platform.openai.com/docs/api-reference/introduction | 
| Ollama               | http://localhost:1143.    | https://olama.com     |
| Amazon Bedrock       |                           | https://aws.amazon.com/bedrock/
| Azure OpenAI         |                           | https://azure.microsoft.com/en-gb/products/ai-foundry/models/openai/ |
| Cloudflar Workers AI |.                          | https://developers.cloudflare.com/workers-ai/|
| Paralon AI           |                           | https://perlonai.com/ |
| Open Router          |                           | https://openrouter.ai |     

## Supported Inference Platforms

| Platfrom       | /v1/completions | /v1/embeddings | /v1/models  | /v1/responses | /v1/images/* | /v1/audio/* | /v1/fine_tuning/* | /v1/assistants |
|----------------|-----------------|----------------|------------|---------------|--------------|-------------|-------------------|----------------|
| Open AI.       | Yes - NB        | Yes            | Yes - NB   | Yes - In Dev  |  Yes - NB    | Yes - NB    | Yes - NB          | Yes - NB       |       
| Ollama         | Yes - NB        | Yes            | Yes - NB.  | Partial - FT  |  NS          |  NS         | NS                | NS             |
| Amazon Bedrock | Yes - NB        | NS             | NS         | Yes In Dev    | NS Y         |  NS         | NS                | NS             |
| Azure OpenAI   | Yes - NB        | ?              | ?          | ?             | ?            | ?           |  ?                |  ?             |      
| Cloudflare. AI | ?               | ?              | ?          | ?             | ?            | ?           |  ?                |  ?             |    
| Paralon AI     | ?               | ?              | ?          | ?             | ?            | ?           |  ?                |  ?             |  
| OpenRouter     | ?               | ?              | ?          | ?             | ?            | ?           |  ?                |  ?             |  

## Supported Vector Database Platforms

| Platfrom             | POST /v1/files    | POST /v1/vector_stores | POST /v1/vector_stores/{vs_id}/files | vector store search | POST /v1/responses with tools: [{type: "file_search"}] |
| ---------------------| ------------------|------------------------| -------------------------------------| --------------------| -------------------------------------------------------| 
| Open AI.             | Yes - NT          |  Yes - NB              | Yes - NB.                            | Yes                 | Yes - NT                                               |  
| Azure AI Search      | ?                 | ?                      | ?                                    | ?                   | ?                                                      |
| Pinecone             | ?                 | ?                      | ?                                    | ?                   | ?                                                      |
| Weaviate             | ?                 | ?                      | ?                                    | ?                   | ?                                                      |
| Qdrant               | ?                 | ?                      | ?                                    | ?                   | ?                                                      |
| Milvus / Zilliz      | ?                 | ?                      | ?                                    | ?                   | ?                                                      |
| Chroma               | ?                 | ?                      | ?                                    | ?                   | ?                                                      |
| Redis (Vector).      | ?                 | ?                      | ?                                    | ?                   | ?                                                      |
| MongoDB Atlas Vector | ?                 | ?                      | ?                                    | ?                   | ?                                                      |

Note: TIBCO AS VDB will require an abstraction layer to support open AI API.  When the abstraction layer is built it will be added to this list for testing. 

## Key to compatibility tables

NS = Platform / OpenAI for Platform does not support this capability
NT = Not Tested.  Not tested, or Testing not complete for this AI platform but capability exists in extension.  
NB = Not Built.  This capability is not currenlty available in the extension but the AI Platform supports it.
NA = Underlying platform does not support this capability. 
FT = Said supported on paper is currently failing testing. 
Partial = Partial capability from platform backend supported. 
? = Unkown till paper execise carried out on platform.

## Ranking Options
These are the ones specifc to Open AI.  There will be a need to add custom values.

Use none when latency matters most
Use default-2024-11-15 when consistency matters most
Use auto when you want OpenAI to manage ranking evolution
