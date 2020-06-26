import React from "react";
import "./App.css";
// import the Container Component from the semantic-ui-react
import { Container } from "semantic-ui-react";
// import the ToDoList component
import Clients from "./clients";

function App() {
  return (
    <div>
      <Container>
        <Clients />
      </Container>
    </div>
  );
}
export default App;