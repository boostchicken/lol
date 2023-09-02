import { lazy, useEffect, useState } from "react";
import History from "./components/History";
import Commands from "./components/Commands";
import ToastContainer from "react-bootstrap/ToastContainer";
import Toast from "react-bootstrap/Toast";
import Container from "react-bootstrap/Container";
const Logo = lazy(() => import("./components/Logo"));
const RS7 = lazy(() => import("./components/Img"));

function App() {
  const [showToast, setShowToast] = useState(false);
  const [toastText, setToastText] = useState("");
  const [showRs, setShowRs] = useState(false);
  useEffect(() => {
    if (Math.random() < 0.1) {
      setShowRs(true);
    }
  }, []);

  useEffect(() => {
    if (toastText !== "") {
      setShowToast(true);
    }
  }, [toastText]);
  return (
    <>
      <div className="App">
        <ToastContainer position="top-end">
          <Toast
            bg="primary"
            show={showToast}
            onClose={() => setShowToast(false)}
            delay={3000}
            autohide
          >
            <Toast.Header>
              <strong className="me-auto">Admin Notification</strong>
            </Toast.Header>
            <Toast.Body>{toastText}</Toast.Body>
          </Toast>
        </ToastContainer>
        <Container>
          {!showRs ? <Logo /> : <RS7 title="A boostchicken in Ascari Blue!" />}
        </Container>
        <Commands toastText={setToastText} />

        <History />
      </div>
    </>
  );
}

export default App;
