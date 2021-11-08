import * as React from 'react'
import './index.css'

const Shop = (props) => {
    return (
        <div className="card m-4 w-50 pb-4">
            <h1 className="card-header">Shop</h1>
            <div className="d-flex justify-content-between flex-wrap">
                {props.products.map((product) => (
                    <div className="col-md-4 mt-4" key={product.id}>
                        <div className="card">
                            <div className="card-img-container">
                                <img
                                    className="card-img-top h-100 w-100"
                                    src={product.image_url}
                                    alt={product.name}
                                />
                            </div>
                            <div className="card-body">
                                <span className="card-title">
                                    <div className="d-flex justify-content-between">
                                        <span>{product.name}</span>
                                        <span>{product.price}$</span>
                                    </div>
                                </span>
                                <div className="d-flex flex-column align-items-end mt-2">
                                    <button
                                        onClick={() =>
                                            props.onAddToCart(product.id)
                                        }
                                        className="btn btn-primary"
                                    >
                                        Add to cart
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    )
}

export { Shop }
