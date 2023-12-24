import FormItem from '@components/FormItem';
import { RouteOnNotOK } from '@lib/Functions';
import { CheckFetchStatus } from '@lib/Status';
import { useQuery } from '@tanstack/react-query';
import { Row, Col } from 'react-bootstrap';
import { useForm } from 'react-hook-form';
import { useNavigate, useSearchParams } from 'react-router-dom';

type sortByType = 'price' | 'stock' | 'sales' | 'relevancy';
type orderType = 'asc' | 'desc';

interface FormProps {
  q: string;
  sortBy: sortByType;
  order: orderType;
  minPrice: string;
  maxPrice: string;
  minStock: string;
  maxStock: string;
  haveCoupon: string;
}

interface IResult {
  products: [
    {
      id: number;
      image_url: string;
      name: string;
      price: number;
      stock: number;
    },
  ];
  shops: [
    {
      id: number;
      image_url: string;
      name: string;
      price: number;
      stock: number;
    },
  ];
}

// const ButtonWrapperStyle = {
//   display: 'flex',
//   justifyContent: 'space-evenly',
// };

// const ButtonStyle = {
//   background: 'var(--button_light)',
//   borderRadius: '10px',
//   width: '50px',
//   height: '50px',
// };

const ProcessInitialValues = (searchParams: URLSearchParams): FormProps => {
  const defaultValues: FormProps = {
    q: searchParams.get('q') ?? '',
    sortBy: (searchParams.get('sortBy') as sortByType) ?? 'relevancy',
    order: (searchParams.get('order') as orderType) ?? 'desc',
    minPrice: '0',
    maxPrice: '0',
    minStock: '0',
    maxStock: '0',
    haveCoupon: 'false',
  };
  const tempMinPrice = Number(searchParams.get('minPrice')) ?? 0;
  const tempMaxPrice = Number(searchParams.get('maxPrice')) ?? 0;
  const tempMinStock = Number(searchParams.get('minStock')) ?? 0;
  const tempMaxStock = Number(searchParams.get('maxStock')) ?? 0;
  defaultValues['minPrice'] = tempMinPrice.toString();
  defaultValues['maxPrice'] =
    tempMinPrice <= tempMaxPrice ? tempMaxPrice.toString() : tempMinPrice.toString();
  defaultValues['minStock'] =
    tempMinStock <= tempMaxStock ? tempMinStock.toString() : tempMaxStock.toString();
  return defaultValues;
};

const Search = () => {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();
  const { register, handleSubmit, getValues } = useForm<FormProps>({
    defaultValues: ProcessInitialValues(searchParams),
  });

  const {
    status,
    data: fetchedData,
    refetch,
  } = useQuery({
    queryKey: ['search'],
    queryFn: async () => {
      const formData = getValues();
      const newSearchParams = new URLSearchParams({
        ...formData,
        sortBy: formData['sortBy'].toString(),
        order: formData['order'].toString(),
      });
      console.log(newSearchParams.toString());
      const resp = await fetch('/api/search?' + newSearchParams.toString(), {
        method: 'GET',
        headers: {
          accept: 'application/json',
        },
      });
      RouteOnNotOK(resp, navigate);
      return await resp.json();
    },
    select: (data) => data as IResult,
    enabled: true,
    refetchOnWindowFocus: false,
  });
  if (status !== 'success') {
    return <CheckFetchStatus status={status} />;
  }
  console.log(fetchedData);

  const onSubmit = async (data: FormProps) => {
    if (data['q'] === '') {
      return;
    }
    refetch();
  };

  return (
    <Row className='flex-wrapper' style={{ paddingTop: '10%' }}>
      <Col md={3} xl={2}>
        <form onSubmit={handleSubmit(onSubmit)}>
          <FormItem label='Search'>
            <input type='text' {...register('q')} />
          </FormItem>
          <FormItem label='SortBy'>
            <select {...register('sortBy')}>
              <option value='relevancy'>Relevancy</option>
              <option value='price'>Price</option>
              <option value='stock'>Stock</option>
              <option value='sales'>Sales</option>
            </select>
          </FormItem>
          <FormItem label='Order'>
            <select {...register('order')}>
              <option value='desc'>Descending</option>
              <option value='asc'>Ascending</option>
            </select>
          </FormItem>
          <FormItem label='Min Price'>
            <input type='number' {...register('minPrice')} />
          </FormItem>
          <FormItem label='Max Price'>
            <input type='number' {...register('maxPrice')} />
          </FormItem>
          <FormItem label='Min Stock'>
            <input type='number' {...register('minStock')} />
          </FormItem>
          <FormItem label='Max Stock'>
            <input type='number' {...register('maxStock')} />
          </FormItem>
          <FormItem label='Have Coupon'>
            <input type='checkbox' {...register('haveCoupon')} />
          </FormItem>
          <button type='submit'>Search</button>
        </form>
      </Col>
    </Row>
  );
};

export default Search;
