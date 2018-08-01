import { types } from "mobx-state-tree";
import { Views } from "../Constants";

const ViewStore = types
  .model({
    currentView: types.enumeration(Object.values(Views))
  })
  .actions(self => ({
    presentView(newView) {
      self.currentView = newView;
    }
  }));

const Defaults = {
  currentView: Views.LOADING
};

const CreateViewStore = overrides =>
  ViewStore.create(Object.assign(Defaults, overrides));

export { ViewStore, CreateViewStore };
