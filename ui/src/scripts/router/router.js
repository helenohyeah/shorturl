import React from "react";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  useParams,
} from "react-router-dom";

const Links = () => {
  return (
    <div>
      <h1>Links</h1>
      <ul>
        <li>
          <Link to="/">Home</Link>
        </li>
        <li>
          <Link to="/u/123">UserId 123</Link>
        </li>
        <li>
          <Link to="/not_found">Not found</Link>
        </li>
        <li>
          <Link to="/login">Login</Link>
        </li>
      </ul>
    </div>
  );
};

const HomePage = () => {
  const { userId } = useParams();
  return (
    <div>
      <Links />
      <h1>{userId ? `User page - userId: ${userId}` : `Home page`}</h1>
    </div>
  );
};

const NotFoundPage = () => {
  return (
    <div>
      <Links />
      <h1>Not found page</h1>
    </div>
  );
};

const LoginPage = () => {
  return (
    <div>
      <Links />
      <h1>Login page</h1>
    </div>
  );
};

const AppRouter = (props) => {
  return (
    <Router>
      <Switch>
        <Route exact path="/" component={HomePage} />
        <Route exact path="/not_found" component={NotFoundPage} />
        <Route exact path="/login" component={LoginPage} />
        <Route path="/u/:userId" component={HomePage} />
        <Route path="*" component={NotFoundPage} />
      </Switch>
    </Router>
  );
};

export default AppRouter;
