'use client';
import { Suspense, useEffect, useState } from "react";
import History from "./components/History";
import Commands from "./components/Commands";
import ToastContainer from "react-bootstrap/ToastContainer";
import Toast from "react-bootstrap/Toast";
import Spinner from 'react-bootstrap/Spinner';
import Container from "react-bootstrap/Container";
import Logo from "./components/Logo";


function App() {
  const [showToast, setShowToast] = useState(false);
  const [toastText, setToastText] = useState("");

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
          <Logo />
        </Container>
        <Suspense fallback={<Spinner animation="border" variant="primary" />}>
          <Commands toastText={setToastText} />
        </Suspense>
        <Suspense fallback={<Spinner animation="border" variant="primary" />}>
           <History />
        </Suspense>
      </div>
    </>
  );
}

export default App;
