import { Component, h } from "preact";
import { Route, Router, RouterOnChangeArgs } from "preact-router";

import data from "../data/hiddenworld.json";
import NetworkHandler from "../handlers/network";
import { tmdbResp } from "../models/interfaces";
import Home from "../routes/home";
import Profile from "../routes/profile";

if ((module as any).hot) {
  // tslint:disable-next-line:no-var-requires
  require("preact/debug");
}

interface AppState {
  loaded: boolean;
}

export default class App extends Component<{}, AppState> {
  public currentUrl?: string;

  public Media: tmdbResp | undefined = undefined;

  public constructor(props: any) {
    super(props);
    this.state = {
      loaded: false
    };
  }
  public handleRoute = (e: RouterOnChangeArgs) => {
    this.currentUrl = e.url;
  };

  public componentDidMount() {
    NetworkHandler.Init().then(id => {
      console.log(id);
      this.Media = data;
      this.setState({
        loaded: true
      });
    });
  }

  public render() {
    if (this.state.loaded) {
      return (
        <div id="app">
          <Router onChange={this.handleRoute}>
            <Route path="/" component={Home} media={this.Media} />
            <Route path="/profile/" component={Profile} user="me" />
            <Route path="/profile/:user" component={Profile} />
          </Router>
        </div>
      );
    }
    return <div id="app">Loading</div>;
  }
}
