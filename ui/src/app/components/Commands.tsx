"use client";
import {
  Suspense,
  useState,
  Dispatch,
  SetStateAction,
  useEffect,
  useDebugValue
} from "react";
import Table from "react-bootstrap/Table";
import Badge from "react-bootstrap/Badge";
import Container from "react-bootstrap/Container";
import Form from "react-bootstrap/Form";
import FloatingLabel from "react-bootstrap/FloatingLabel";
import Button from "react-bootstrap/Button";
import Link from "next/link";
import *  as api from "@boostchicken/lol-api";

interface CommandProps {
  toastText: Dispatch<SetStateAction<string>>;
}

function Commands(props: CommandProps) {
  const toastText = props.toastText;

  const [newCommand, setNewCommand] = useState("");
  const [newType, setNewType] = useState("Alias");
  const [newValue, setNewValue] = useState("");
  const [delValue, setDelValue] = useState("");

  const { data: config, mutate, error } =api.useGetLiveConfig();
  const { trigger: addCmd } = api.useAddCommand(newCommand, getTypeValue(), {
    url: newValue,
  });
  const { trigger: deleteCmd } = api.useDeleteCommand(delValue);

  function getTypeValue() {
    switch (newType) {
      case "Alias":
        return api.addCommandPathParamsType.Alias;
      case "Redirect":
        return api.addCommandPathParamsType.Redirect;
      case "RedirectVarArgs":
        return api.addCommandPathParamsType.RedirectVarArgs;
    }
    return api.addCommandPathParamsType.Alias;
  }

  useEffect(() => {
    async function fetch() {
      return mutate(deleteCmd());
    }
    if(delValue === "") return
    fetch()
      .then(() => {
        toastText(`Deleted ${delValue}`);
      })
      .catch((err) => {
        toastText(`error ${err}`);
      });
  }, [deleteCmd, mutate,delValue,toastText]);
  
  const addEntry = () => {
    async function fetch() {
      return mutate(addCmd());
    }
    fetch()
      .then(() => {
        toastText(`Added ${newCommand}`);
      })
      .catch(async (err) => {
        toastText(`error ${err}`);
      });
  };

  if (error) return <div>Error</div>;

  return (
    <Container>
      <h3>
        Commands <Badge bg="danger"> {config?.Entries?.length} </Badge>
        <Link href="/api">
          <Button variant="danger">API Docs</Button>
        </Link>
      </h3>
        <Table striped bordered hover variant="dark">
          <thead>
            <tr>
              <th>Command</th>
              <th>Mode</th>
              <th>URL Template</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            <tr key="add">
              <td>
                <FloatingLabel controlId="floatingCommand" label="Command">
                  <Form.Control
                    inputMode="text"
                    placeholder="Command"
                    value={newCommand}
                    onChange={(e) => setNewCommand(e.currentTarget.value)}
                  />
                </FloatingLabel>
              </td>
              <td>
                <FloatingLabel controlId="floatingSelect" label="Mode">
                  <Form.Select
                    aria-label="Mode"
                    onChange={(e) => setNewType(e.currentTarget.value)}
                  >
                    <option value="Alias">Alias</option>
                    <option value="Redirect">Redirect</option>
                    <option value="RedirectVarArgs">VarArgs</option>
                  </Form.Select>
                </FloatingLabel>
              </td>
              <td>
                <FloatingLabel controlId="floatingValue" label="URL Template">
                  <Form.Control
                    type="text"
                    inputMode="url"
                    value={newValue}
                    placeholder="URL"
                    onChange={(e) => setNewValue(e.currentTarget.value)}
                  />
                </FloatingLabel>
              </td>
              <td>
                <Button variant="primary" type="button" onClick={addEntry}>
                  Add
                </Button>
              </td>
            </tr>
            {config?.Entries.map((item, idx) => (
              <tr key={idx}>
                <td>{item.Command}</td>
                <td>{item.Type}</td>
                <td>{item.Value}</td>
                <td>
                  <Button
                    variant="danger"
                    type="button"
                    onClick={() => {
                      setDelValue(`${item.Command}`);
                    }}
                  >
                    Delete
                  </Button>
                </td>
              </tr>
            ))}
          </tbody>
        </Table>
    </Container>
  );
}

export default Commands;
