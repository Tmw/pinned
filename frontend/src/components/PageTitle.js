import React from "react";
import PropTypes from "prop-types";
import { Helmet } from "react-helmet";

const mainTitleDefault = "Pinned! - ";
const makeTitle = ({ mainTitle = mainTitleDefault, subtitle }) =>
  mainTitle.concat(subtitle);

const PageTitle = ({ subtitle }) => <Helmet title={makeTitle({ subtitle })} />;

PageTitle.propTypes = {
  subtitle: PropTypes.string.isRequired
};

export default PageTitle;
