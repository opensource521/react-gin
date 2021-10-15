import * as redux from 'react-redux'
import App from './App'
import { mount, ReactWrapper } from 'enzyme'
import InputForm from './components/InputForm'

describe('<App />', () => {
  const smiles = 'CCC(CC)CCC'

  it('should display smiles in store', async () => {
    const useSelectorSpy = jest.spyOn(redux, 'useSelector')
    const useDispatchSpy = jest.spyOn(redux, 'useDispatch')
    const mockDispatchFn = jest.fn()
    useSelectorSpy.mockReturnValueOnce(smiles).mockReturnValueOnce('')
    useDispatchSpy.mockReturnValue(mockDispatchFn)

    const wrapper: ReactWrapper = mount(<App />)

    expect(wrapper.find(InputForm).props()).toHaveProperty('smiles', smiles)
  })

  it('should dispatch action when the form submits', () => {
    const useSelectorSpy = jest.spyOn(redux, 'useSelector')
    const useDispatchSpy = jest.spyOn(redux, 'useDispatch')
    const mockDispatchFn = jest.fn()
    useSelectorSpy.mockReturnValueOnce({}).mockReturnValueOnce('')
    useDispatchSpy.mockReturnValue(mockDispatchFn)

    const wrapper: ReactWrapper = mount(<App />)
    wrapper.find(InputForm).props().onSubmit(smiles)

    expect(mockDispatchFn).toHaveBeenCalled()

    useDispatchSpy.mockClear()
  })
})
