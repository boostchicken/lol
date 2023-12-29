'use client';
import SwaggerUI from "swagger-ui-react";
import "swagger-ui-react/swagger-ui.css";
import "./SwaggerDark.css"

function Api() {
  return (
    <SwaggerUI
      tryItOutEnabled={true}
      requestSnippetsEnabled={true}
      url="./openapi.yaml"
    />
  );
}

export default Api;
