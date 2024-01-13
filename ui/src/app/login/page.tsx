import Logo from "../components/Logo";
import Container from "react-bootstrap/Container";
import Button from "react-bootstrap/Button";
import Stack from "react-bootstrap/Stack";
import Link from "next/link";

function Login() {

    return (
      <Container>
            <Logo />

            <Stack gap={2} className="col-md-5 mx-auto">
                <Button variant="primary">Login</Button>
              <Link  href={"/register"}>
                <Button variant="danger">Register</Button>
            </Link>
            </Stack>
        </Container>
    )

}

export default Login;