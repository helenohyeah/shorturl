import { bindActionCreators } from 'redux';
import { useDispatch } from 'react-redux';
import { useMemo } from 'react';

// Taken from https://react-redux.js.org/next/api/hooks#recipe-useactions
// to reduce boiler-plate when providing actions to components
function useActions(actions, deps) {
    const dispatch = useDispatch();

    return useMemo(
        () => {
            if (Array.isArray(actions)) {
                return actions.map(a => bindActionCreators(a, dispatch));
            }
            return bindActionCreators(actions, dispatch);
        },
        deps ? [dispatch, ...deps] : [dispatch] //eslint-disable-line
    );
}
export default useActions;
