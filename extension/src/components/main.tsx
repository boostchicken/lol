
import { useState, useEffect} from "react"
import {useGetLiveConfig} from '@boostchicken/lol-api'
import Button from "react-bootstrap/Button"
import Container from "react-bootstrap/Container"
import Form from "react-bootstrap/Form"
import FloatingLabel from "react-bootstrap/FloatingLabel"
function Main() {
    const [conf, setConf] = useState<any>("")

    useEffect(() => {
      const { data: resp, mutate, error } = useGetLiveConfig()
      setConf(resp)
    }, [])
    return(

<Container>
    Command Count: {conf.Entries?.length}
    <Form>
    <FloatingLabel controlId="floatingPassword" label="API Key">
        <Form.Control value={conf.Bind} type="password" placeholder="Api" />
    </FloatingLabel>
    <Button variant="primary" >Set</Button> 
    <Button variant="danger">Clear </Button>
    <Form.Check // prettier-ignore
        type="switch"
        id="custom-switch"
        label="Enable Config Cache"
      />
    </Form>
    </Container>

)
}
export default Main