import { Component, h } from "preact";
import { Route, Router, RouterOnChangeArgs } from "preact-router";

import { format } from "url";
import NetworkHandler from "../handlers/network";
import Home from "../routes/home";
import Profile from "../routes/profile";

if ((module as any).hot) {
  // tslint:disable-next-line:no-var-requires
  require("preact/debug");
}

export default class App extends Component {
  public currentUrl?: string;
  public handleRoute = (e: RouterOnChangeArgs) => {
    this.currentUrl = e.url;
  };

  public componentDidMount() {
    NetworkHandler.Init().then(id => {
      console.log(id);
    });
  }

  public render() {
    return (
      <div id="app">
        <Router onChange={this.handleRoute}>
          <Route path="/" component={Home} />
          <Route path="/profile/" component={Profile} user="me" />
          <Route path="/profile/:user" component={Profile} />
        </Router>
      </div>
    );
  }
}
