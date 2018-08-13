import { types } from "mobx-state-tree";
import { Views } from "Constants";

const ViewStore = types
  .model({
    currentView: types.enumeration(Object.values(Views)),
    error: types.maybeNull(types.string)
  })
  .actions(self => ({
    presentView(newView, error = null) {
      self.currentView = newView;
      self.error = error;
    }
  }));

const Defaults = {
  currentView: Views.LOADING
};

const CreateViewStore = overrides =>
  ViewStore.create(Object.assign(Defaults, overrides));

export { ViewStore, CreateViewStore };
