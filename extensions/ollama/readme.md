Flogo Ollama Client Extension
=============================

Please do not use this for your own projects at this stage as it is still in very early stage of development.  If you are a TIBCO customer and would like to use the extension or have feeback then please reach out to the repository owner or your account team. 

Overview
--------
The Flogo ollama extesion acts as a client for the APIs exposed from ollama application.  It is in the very early days of development and there is a focus on supporting text generation, image generation, embedings creation and storage using the Responses API. 

Mini Roadmap 
------------

Focus on Response, Chat and Embeddings first.  
Then look at model managemnt capabilites for bootstrapping and cleanup.   For example load a model if it is not available.   Load into memory at start of app etc with keep alive.

OC Endpoints
-------------
 
Generate a response
Generate a chat completion
Generate Embeddings

---------------------
Some of these maybe required for bootstrapping / clean up ...   

List Models
List Running Models
Show model details
Create a model
Copy a model
Pull a model
Push a model
delete a model
get version


OC Model Management Endpoints
------------------------------

Create a Model
List Local Models
Show Model Information
Copy a Model
Delete a Model
Pull a Model
Push a Model
List Running Models
Version


| Client Activity                     | Status                        | Go API Library Status  |
| ------------------------------------| ------------------------------| ------------------------------|
|                                     | In Dev for text and Images    | v0.11.4                       |
| Images API                          | In Dev for text and Images    | v                             |
| Embedings API                       | In Dev for text and Images    | v1.12.0                       | 

Ollamam library version

v0.11.4

Ollama Use Cases 
-----------------

| Client Activity                     | Status                        |
| ----------- ----------------------- | ------------------------------|
| Text Generation                     | In Dev                        |
| Image Generation                    | In Dev                        |
| Audio Generation                    | Currently out of scope        |
| Deep Research                       | Currently out of scope        |  
| Embeddings                          | In Dev                        |
| Moderation                          | Currently out of scope        |
| Model Managment                     | Out of scope                  | 

Flogo Example Use Caes 
----------------------

These are some of the example use cases we are being reviewed as part of the client extension development.

Use Flogo to get Structured Data from a model and store it in Mongo DB.

Use Flogo to generate Text and Images and then forward and send to a message broker.  

Use Flogo to Analyse images that are part of an event stream

Use flogo to first expose APIs with MCP and then use Remote MCP with an LLM.

Use flogo to do a file search

Use flogo to do a web search


Ollama Documentation 
--------------------
Official Ollama documentation on how to use Ollama and details about the API used for this connector can be found at this URL

https://github.com/ollama/ollama/blob/main/docs/README.md

More Information
----------------

For more information please reach out to the repository owner.


Embedding Prefix and Body
-------------------------

When working with some models it is important to add perfix and body to the text before passing to the embedding model, below are examples of how to achieve that with the current activity

"embeddingText": "=string.concat(\"title:\",$flow.output.fileMetadata.name,\" | text: \",coerce.toString($activity[Generatecompletion].outputText

"embeddingText": "=string.concat(\"task: search result | query: \", $flow.body.searchString)