import React, { Component } from "react";
import axios from "axios";
import "./App.css";
import { Card, Container, Image, Button, Icon } from "semantic-ui-react";

let endpoint = "http://localhost:8080";

class LandingPage extends Component {
  render() {
    return (
      <Container>
        <div className="App">
          <h1>Abe Legal</h1>
          <Card.Group>
            <Card>
              <Image
                src="https://react.semantic-ui.com/images/avatar/large/matthew.png"
                wrapped
                ui={false}
              />
              <Card.Content>
                <Card.Header>For Lawyers</Card.Header>
                <Card.Meta>
                  <span className="date">Joined in 2015</span>
                </Card.Meta>
                <Card.Description>
                  <Button primary href="/lawyerdashboard/sign_in">
                    Click Here!
                  </Button>
                </Card.Description>
              </Card.Content>
              <Card.Content extra>
                <a>
                  <Icon name="user" />
                  Current Users:
                </a>
              </Card.Content>
            </Card>
            <Card>
              <Image
                src="https://react.semantic-ui.com/images/avatar/large/matthew.png"
                wrapped
                ui={false}
              />
              <Card.Content>
                <Card.Header>For Clients</Card.Header>
                <Card.Meta>
                  <span className="date">Joined in 2015</span>
                </Card.Meta>
                <Card.Description>
                  <Button primary href="/client">
                    Click Here!
                  </Button>
                </Card.Description>
              </Card.Content>
              <Card.Content extra>
                <a>
                  <Icon name="user" />
                  Current Users:
                </a>
              </Card.Content>
            </Card>
          </Card.Group>
        </div>
      </Container>
    );
  }
}

export default LandingPage;
