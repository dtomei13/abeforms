import React from "react";
import "./App.css";
// import the Container Component from the semantic-ui-react
import { Container } from "semantic-ui-react";
// import the ToDoList component
import Clients from "./clients";
import LawyerDashboard from "./LawyerDashboard";
import LawyerSignUp from "./LawyerSignUp"
import LawyerSignIn from "./LawyerSignIn"

import {
  BrowserRouter as Router,
  Route,
  Switch,
  Link,
  Redirect
} from "react-router-dom"

function App() {
  return (
    <Router>
    <Switch>

        <Route exact path="/client" component = {Clients} />
        <Route exact path="/lawyerdashboard/sign_in" component = {LawyerSignIn} />
        <Route exact path="/lawyerdashboard/sign_up" component = {LawyerSignUp} />
        <Route exact path="/lawyerdashboard" component = {LawyerSignIn} />

      </Switch>
      </Router>
  );
}
export default App;