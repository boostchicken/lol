import Logo from "../components/Logo";
import Container from "react-bootstrap/Container";
import Button from "react-bootstrap/Button";
import Stack from "react-bootstrap/Stack";
import  FormText from "react-bootstrap/FormText";

function Register() {

    return (
      <Container>
            <Logo />

            <Stack gap={2} className="col-md-5 mx-auto">
                <FormText  placeholder="email" aria-label="email" />
                <Button variant="danger">Register</Button>
               
            </Stack>
                   </Container>
    )

}

export default Register;