
import { RedocStandalone } from 'redoc';
function ApiDocs() {
  return (
    <RedocStandalone specUrl='assets/openapi.yaml' />  
  );
}

export default ApiDocs;