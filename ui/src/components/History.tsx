"use client";
import Table from "react-bootstrap/Table";
import Badge from "react-bootstrap/Badge";
import Container from "react-bootstrap/Container";
import Spinner from "react-bootstrap/Spinner";

import { useGetHistory } from "../gen/hooks/useGetHistory";

function History() {
  const { data: history, error, isLoading } = useGetHistory();
  if (error) return <div>Error</div>;
  if (isLoading) return <Spinner animation="border" variant="primary" />;
  return (
    <Container>
      <h3>
        History <Badge bg="primary">{history?.length}</Badge>
      </h3>
      <Table striped bordered hover variant="dark">
        <thead>
          <tr>
            <th>Command</th>
            <th>Result</th>
          </tr>
        </thead>
        <tbody>
          {history?.map((item, idx) => (
            <tr key={idx}>
              <td>{item.Command}</td>
              <td>
                <a href={item.Result}>{item.Result}</a>
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
    </Container>
  );
}

export default History;
